package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/handler"
	"github.com/rusrom/yt-todo/pkg/repository"
	"github.com/rusrom/yt-todo/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading environment variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config %s", err.Error())
	}

	db, err := repository.NewDb(repository.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("fail to initialize database: %s", err.Error())
	}

	repos := repository.NewTodoRepository(db)
	services := service.NewTodoService(repos)
	handlers := handler.NewTodoHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error while running web server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
