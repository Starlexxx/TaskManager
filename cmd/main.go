package main

import (
	"TaskManager"
	"TaskManager/pkg/handler"
	"log"
)

func main() {
	serverHandler := new(handler.Handler)

	server := new(TaskManager.Server)
	if err := server.Start("8080", serverHandler.InitRoutes()); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}
