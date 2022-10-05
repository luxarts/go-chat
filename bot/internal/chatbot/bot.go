package chatbot

import (
	"bot/internal/defines"
	"bot/internal/domain"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"sync"
)

type Handler func(ctx *Context) string

type Bot interface {
	OnCommand(command string, handler Handler)
	Run() error
}

type bot struct {
	backendURL string
	commands   map[string]Handler
	mu         sync.RWMutex
	conn       *websocket.Conn
}

func New(backendURL string) Bot {
	conn, _, err := websocket.DefaultDialer.Dial("ws://"+backendURL+defines.APIBackendPathChatroom, nil)
	if err != nil {
		log.Fatalf("failed to connect ws: %v", err)
	}

	return &bot{
		commands:   make(map[string]Handler),
		backendURL: backendURL,
		conn:       conn,
	}
}

func (b *bot) OnCommand(command string, handler Handler) {
	b.commands[command] = handler
}

func (b *bot) Run() error {
	for {
		msgType, dataBytes, err := b.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("unexpected error on close: %v\n", err)
			}
		}

		if msgType == websocket.TextMessage {
			var data domain.Data

			if err := json.Unmarshal(dataBytes, &data); err != nil {
				log.Printf("error unmarhaling: %v\n", err)
				return err
			}

			for cmd, handler := range b.commands {
				if strings.Index(data.Msg, cmd) == 0 {
					ctx := &Context{Payload: data.Msg}
					response := handler(ctx)

					err := b.conn.WriteJSON(domain.Data{
						User: "Stock Bot",
						Msg:  response,
					})
					if err != nil {
						log.Printf("error responding: %v\n", err)
					}

					break
				}
			}
		}
	}
}
