package cmd

import (
	"github.com/Likhon22/ecom/config"
	"github.com/Likhon22/ecom/rest"
	"github.com/Likhon22/ecom/rest/handlers/product"
	"github.com/Likhon22/ecom/rest/handlers/user"
)

func Serve() {

	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	server := rest.NewServer(productHandler, userHandler)

	server.StartServer(&cnf)
}
