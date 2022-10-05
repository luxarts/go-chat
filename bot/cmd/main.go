package main

import (
	"bot/internal/chatbot"
	"bot/internal/controller"
	"bot/internal/defines"
	"bot/internal/repository"
	"bot/internal/service"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

func main() {
	backendURL := os.Getenv(defines.EnvBackendURL)
	if backendURL == "" {
		log.Fatalln(fmt.Sprintf("Env var %s is empty.", defines.EnvBackendURL))
	}

	bot := chatbot.New(backendURL)

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
