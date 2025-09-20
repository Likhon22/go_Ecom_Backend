package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var config *Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
	SecretKey   string
	DBUrl       string
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
	dbUrl := os.Getenv("DB_URL")
	config = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		SecretKey:   secretKey,
		DBUrl:       dbUrl,
	}
	if config.DBUrl == "" {
		fmt.Println("Missing required environment variables: DB_URL")
		os.Exit(1)

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
func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}
