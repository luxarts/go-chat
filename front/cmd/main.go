package main

import (
	"frontend/internal/router"
	"log"
)

func main() {
	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
