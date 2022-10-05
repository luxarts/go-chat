package main

import (
	"fmt"
	"frontend/internal/defines"
	"frontend/internal/router"
	"log"
	"os"
)

func main() {
	backendURL := os.Getenv(defines.EnvBackendURL)
	if backendURL == "" {
		log.Fatalln(fmt.Sprintf("Env var %s is empty.", defines.EnvBackendURL))
	}

	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
