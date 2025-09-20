package cmd

import (
	"fmt"
	"os"

	"github.com/Likhon22/ecom/config"
	"github.com/Likhon22/ecom/infra/db"
	"github.com/Likhon22/ecom/repo"
	"github.com/Likhon22/ecom/rest"
	"github.com/Likhon22/ecom/rest/handlers/product"

	"github.com/Likhon22/ecom/rest/handlers/user"
	"github.com/Likhon22/ecom/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()
	db, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	middleware := middleware.NewMiddlewares(cnf)
	userRepo := repo.NewUserRepo(db)
	productRepo := repo.NewProductRepo(db)
	productHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(userRepo)

	server := rest.NewServer(productHandler, userHandler, cnf)
	server.StartServer()

}
