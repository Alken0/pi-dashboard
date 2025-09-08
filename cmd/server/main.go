package main

import (
	"log"
	"pi-dashboard/internal/config"
	"pi-dashboard/internal/server"
)

func main() {
	cfg := config.Load()

	if err := server.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
