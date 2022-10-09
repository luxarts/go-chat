package wsserver

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WSServer interface {
	CreateConnection(w http.ResponseWriter, r *http.Request) *connection
	Subscribe(conn *connection)
	SendMessage(conn *connection, d *domain.Data)
	run()
	registerUser(s *subscription)
	unregisterUser(s *subscription)
	broadcastMessage(d *domain.Data)
}

type wsServer struct {
	connections    map[*connection]bool
	dataChan       chan domain.Data
	registerChan   chan subscription
	unregisterChan chan subscription
	upgrader       websocket.Upgrader
	config         *Config
	msgRepo        repository.MessagesRepository
}

type Config struct {
	ReadBufferSize  int
	WriteBufferSize int
}

func New(config *Config, msgRepo repository.MessagesRepository) WSServer {
	u := websocket.Upgrader{
		ReadBufferSize:  config.ReadBufferSize,
		WriteBufferSize: config.ReadBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	wss := &wsServer{
		connections:    make(map[*connection]bool),
		dataChan:       make(chan domain.Data),
		registerChan:   make(chan subscription),
		unregisterChan: make(chan subscription),
		upgrader:       u,
		config:         config,
		msgRepo:        msgRepo,
	}

	wss.run()

	return wss
}
func (wss *wsServer) run() {
	go func() {
		for {
			select {
			case s := <-wss.registerChan:
				wss.registerUser(&s)

			case s := <-wss.unregisterChan:
				wss.unregisterUser(&s)

			case d := <-wss.dataChan:
				wss.broadcastMessage(&d)
				wss.msgRepo.Add(&d)
			}
		}
	}()
}
func (wss *wsServer) registerUser(s *subscription) {
	if wss.connections == nil {
		wss.connections = make(map[*connection]bool)
	}
	wss.connections[s.conn] = true
}
func (wss *wsServer) unregisterUser(s *subscription) {
	if wss.connections != nil {
		if _, ok := wss.connections[s.conn]; ok {
			delete(wss.connections, s.conn)
			close(s.conn.send)
		}
	}
}
func (wss *wsServer) broadcastMessage(d *domain.Data) {
	for c := range wss.connections {
		dataBytes, err := json.Marshal(d)
		if err != nil {
			log.Printf("error marshaling: %v\n", err)
			break
		}

		select {
		case c.send <- dataBytes:
		default:
			close(c.send)
			delete(wss.connections, c)
		}
	}
}
func (wss *wsServer) SendMessage(c *connection, d *domain.Data) {
	dataBytes, err := json.Marshal(d)
	if err != nil {
		log.Printf("error marshaling: %v\n", err)
		return
	}

	select {
	case c.send <- dataBytes:
	default:
		close(c.send)
		delete(wss.connections, c)
	}
}

func (wss *wsServer) CreateConnection(w http.ResponseWriter, r *http.Request) *connection {
	wsc, err := wss.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return &connection{
		send: make(chan []byte, wss.config.WriteBufferSize),
		wsc:  wsc,
	}
}
func (wss *wsServer) Subscribe(conn *connection) {
	s := subscription{
		conn: conn,
	}

	wss.registerChan <- s

	s.writeRoutine()
	s.readRoutine(wss.unregisterChan, wss.dataChan)
}
