package cmd

import (
	"github.com/Likhon22/ecom/config"
	"github.com/Likhon22/ecom/rest"
)

func Serve() {

	cnf := config.GetConfig()

	rest.StartServer(&cnf)
}
