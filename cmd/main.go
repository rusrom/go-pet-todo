package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/handler"
	"github.com/rusrom/yt-todo/pkg/repository"
	"github.com/rusrom/yt-todo/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Todo App API
// @version 1.0.0
// @description API Server for TodoList Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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

	// Gracefull Shutdown block start
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error while running web server: %s", err.Error())
		}
	}()

	logrus.Info("Todo app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Info("Todo app shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error ocurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error ocurred on db connection close: %s", err.Error())
	}

	// Gracefull Shutdown block end
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
