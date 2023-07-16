package main

import (
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/handler"
	"github.com/rusrom/yt-todo/pkg/repository"
	"github.com/rusrom/yt-todo/pkg/service"
	"log"
)

func main() {
	repos := repository.NewTodoRepository()
	services := service.NewTodoService(repos)
	handlers := handler.NewTodoHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running web server: %s", err.Error())
	}
}
