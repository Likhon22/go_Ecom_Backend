package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var config Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	config = Config{
		Version:     os.Getenv("VERSION"),
		ServiceName: os.Getenv("SERVICE_NAME"),
		HttpPort:    os.Getenv("HTTP_PORT"),
	}
}
func GetConfig() Config {
	loadConfig()
	if config.HttpPort == "" || config.ServiceName == "" || config.Version == "" {
		fmt.Println("Missing required environment variables")
		os.Exit(1)

	}
	return config
}
