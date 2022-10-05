package main

import (
	"bot/internal/chatbot"
	"bot/internal/controller"
	"bot/internal/defines"
	"bot/internal/repository"
	"bot/internal/service"
	"github.com/go-resty/resty/v2"
	"log"
)

func main() {
	bot := chatbot.New(defines.APIBackendURL)

	mapCommands(bot)

	if err := bot.Run(); err != nil {
		log.Fatalln(err)
	}
}

func mapCommands(b chatbot.Bot) {
	// RestClient
	rc := resty.New()

	// Repositories
	repo := repository.NewStooqRepository(rc)

	// Services
	svc := service.NewQuoteRepository(repo)

	// Controllers
	ctrl := controller.NewCommandsController(svc)

	// Commands
	b.OnCommand(defines.CommandStock, ctrl.Stock)
}
