package main

import (
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.TodoHandler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running web server: %s", err.Error())
	}
}
