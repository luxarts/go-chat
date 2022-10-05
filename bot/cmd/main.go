package main

import (
	"bot/internal/chatbot"
	"bot/internal/controller"
	"bot/internal/defines"
	"bot/internal/repository"
	"bot/internal/service"
	"log"
)

func main() {
	backendURL := "localhost:9090"
	bot := chatbot.New(backendURL)

	mapCommands(bot)

	bot.SetMessage("/stock=aapl.us")

	if err := bot.Run(); err != nil {
		log.Fatalln(err)
	}
}

func mapCommands(b chatbot.Bot) {
	// Repositories
	repo := repository.NewStooqRepository()

	// Services
	svc := service.NewQuoteRepository(repo)

	// Controllers
	ctrl := controller.NewCommandsController(svc)

	// Commands
	b.OnCommand(defines.CommandStock, ctrl.Stock)
}
