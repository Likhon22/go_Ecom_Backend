package main

import (
	"log"

	"github.com/Likhon22/ecom/cmd"
	"github.com/Likhon22/ecom/config"
	_ "github.com/Likhon22/ecom/database"
)

func main() {
	config := config.GetConfig()
	log.Printf("Starting %s version %s on port %s", config.ServiceName, config.Version, config.HttpPort)
	cmd.Serve()

}
