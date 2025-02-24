package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitCfg() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func GetPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = viper.GetString("server.port")
		log.Printf("SERVER_PORT not set, using default %s", port)
	}
	return port
}
