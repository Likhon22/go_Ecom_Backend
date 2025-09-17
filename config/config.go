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
	SecretKey   string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")
	secretKey := os.Getenv("SECRET_KEY")
	config = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		SecretKey:   secretKey,
	}
	if config.HttpPort == "" || config.ServiceName == "" || config.Version == "" {
		fmt.Println("Missing required environment variables", config.HttpPort, config.ServiceName, config.Version)
		os.Exit(1)

	}
	if config.SecretKey == "" {
		fmt.Println("Missing required environment variables: SECRET_KEY")
		os.Exit(1)
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
