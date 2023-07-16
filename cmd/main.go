package main

import (
	todo "github.com/rusrom/yt-todo"
	"log"
)

func main() {
	srv := new(todo.Server)

	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error while running web server: %s", err.Error())
	}
}
