package chatbot

import (
	"strings"
	"sync"
)

type Handler func(ctx *Context)

type Bot interface {
	OnCommand(command string, handler Handler)
	Run() error
	SetMessage(msg string)
}

type bot struct {
	backendURL string
	commands   map[string]Handler
	mu         sync.RWMutex
	msg        string
}

func New(backendURL string) Bot {
	return &bot{
		commands:   make(map[string]Handler),
		backendURL: backendURL,
	}
}

func (b *bot) OnCommand(command string, handler Handler) {
	b.commands[command] = handler
}

func (b *bot) Run() error {
	if b.messageReceived() {
		for cmd, handler := range b.commands {
			if strings.Index(b.msg, cmd) == 0 {
				ctx := &Context{Payload: b.msg}
				handler(ctx)
				break
			}
		}
	}

	return nil
}

func (b *bot) messageReceived() bool {
	return true
}

func (b *bot) SetMessage(msg string) {
	b.msg = msg
}
