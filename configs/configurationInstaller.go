package configs

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("failed to load .env file %s", err.Error())
	}
}

// функция для добавление конфигов
func initConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("failed to load config file %s", err.Error())
	}
}

func ConfigurationInstaller() {
	initConfig()
	loadEnv()
}
