package controller

import (
	"bot/internal/chatbot"
	"bot/internal/service"
	"fmt"
	"strings"
)

type CommandsController interface {
	Stock(ctx *chatbot.Context) string
}

type commandsController struct {
	quoteSvc service.QuoteService
}

func NewCommandsController(quoteSvc service.QuoteService) CommandsController {
	return &commandsController{quoteSvc: quoteSvc}
}

func (c *commandsController) Stock(ctx *chatbot.Context) string {
	payload := ctx.Payload
	stockCode := payload[strings.Index(payload, "=")+1:]

	quote, err := c.quoteSvc.GetQuote(stockCode)
	if err != nil {
		return fmt.Sprintf("An error ocurred: %v", err)
	}

	return fmt.Sprintf("%s quote is $%.2f per share.", quote.Symbol, quote.Close)
}
