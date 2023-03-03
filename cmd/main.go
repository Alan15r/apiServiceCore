package main

import (
	"fmt"
	"github.com/Alan15r/apiServiceCore/configs"
	"github.com/Alan15r/apiServiceCore/pkg/handler"
	"github.com/Alan15r/apiServiceCore/pkg/repository"
	"github.com/Alan15r/apiServiceCore/pkg/server"
	"github.com/Alan15r/apiServiceCore/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	dbPool, err := repository.NewConnPool(&config.Database)
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(dbPool)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	logrus.Println("Запуск сервера")
	if err := srv.Run(fmt.Sprintf(":%d", config.Port), handlers.InitRoutes().Handler); err != nil {
		logrus.Fatal(err)
	}
}
