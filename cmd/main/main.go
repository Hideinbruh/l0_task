package main

import (
	cacheOrder "awesomeProject/cache"
	handlerOrder "awesomeProject/internal/handler"
	"awesomeProject/internal/model"
	"awesomeProject/internal/nats"
	"awesomeProject/internal/repository"
	serverOrder "awesomeProject/internal/server"
	"awesomeProject/internal/service"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var order *model.Order
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logger.Fatalf("Error initializing config: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Database: viper.GetString("db.database"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		SslMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logger.Fatalf("Error to connect postgres: %s", err)
	}

	cache := cacheOrder.NewCache()
	server := new(serverOrder.Server)
	repo := repository.NewRepository(db)
	cacheRepo := repository.NewOrderCache(cache)
	services := service.NewService(repo, cacheRepo)
	handler := handlerOrder.NewHanlder(services, cache)

	subscriber := nats.NewSubscriber("test-cluster", "client1", "nats://localhost:4222", "example",
		services)
	subscriber.ConnectAndSubscribe(cache, services, order)
	if err != nil {
		logger.Errorf("Error connecting and subscribing to Nats-streaming: %s", err)
	}
	go func() {
		if err := server.Run("localhost:8080", handler.InitRoutes()); err != nil {
			log.Fatalf("Run server error: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	err = db.Close()
	if err != nil {
		logger.Errorf("Error closing db: %s", err)
	}

	err = server.Shutdown(context.Background())
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println("App stopped")
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
