package main

import (
	_ "github.com/lib/pq"
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/handler"
	"github.com/rusrom/yt-todo/pkg/repository"
	"github.com/rusrom/yt-todo/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization config %s", err.Error())
	}

	db, err := repository.NewDb(repository.DbConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("fail to initialize database: %s", err.Error())
	}

	repos := repository.NewTodoRepository(db)
	services := service.NewTodoService(repos)
	handlers := handler.NewTodoHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running web server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
