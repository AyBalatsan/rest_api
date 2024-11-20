package main

import (
	"os"

	todo "github.com/AyBalatsan/Rest_API"
	"github.com/AyBalatsan/Rest_API/configs"
	"github.com/AyBalatsan/Rest_API/pkg/handler"
	"github.com/AyBalatsan/Rest_API/pkg/repository"
	"github.com/AyBalatsan/Rest_API/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Устанавливаем все зависимости
	configs.ConfigurationInstaller()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		DBName:   viper.GetString("database.dbname"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  viper.GetString("database.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error init DB %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		logrus.Printf("error occured while running http server: %s", err.Error())
	}
}
