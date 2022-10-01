package wsserver

import (
	"backend/internal/domain"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WSServer interface {
	CreateConnection(w http.ResponseWriter, r *http.Request) *connection
	Subscribe(conn *connection)
	run()
	registerUser(s *subscription)
	unregisterUser(s *subscription)
	sendMessage(m *domain.Message)
}

type wsServer struct {
	connections    map[*connection]bool
	messageChan    chan domain.Message
	registerChan   chan subscription
	unregisterChan chan subscription
	upgrader       websocket.Upgrader
	config         *Config
}

type Config struct {
	ReadBufferSize  int
	WriteBufferSize int
}

func New(config *Config) WSServer {
	u := websocket.Upgrader{
		ReadBufferSize:  config.ReadBufferSize,
		WriteBufferSize: config.ReadBufferSize,
	}

	wss := &wsServer{
		connections:    make(map[*connection]bool),
		messageChan:    make(chan domain.Message),
		registerChan:   make(chan subscription),
		unregisterChan: make(chan subscription),
		upgrader:       u,
		config:         config,
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

			case m := <-wss.messageChan:
				wss.sendMessage(&m)
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
func (wss *wsServer) sendMessage(m *domain.Message) {
	for c := range wss.connections {
		select {
		case c.send <- m.Data:
		default:
			close(c.send)
			delete(wss.connections, c)
		}
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
	s.readRoutine(wss.unregisterChan, wss.messageChan)
}
