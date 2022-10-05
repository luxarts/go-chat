package chatbot

import "fmt"

type Context struct {
	Payload string
}

func (ctx *Context) Respond(msg string) {
	fmt.Printf("Bot: %v\n", msg)
}
