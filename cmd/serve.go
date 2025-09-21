package cmd

import (
	"fmt"
	"os"

	"github.com/Likhon22/ecom/config"
	"github.com/Likhon22/ecom/infra/db"
	"github.com/Likhon22/ecom/product"
	"github.com/Likhon22/ecom/repo"
	"github.com/Likhon22/ecom/rest"
	productHandler "github.com/Likhon22/ecom/rest/handlers/product"
	"github.com/Likhon22/ecom/user"

	userHandler "github.com/Likhon22/ecom/rest/handlers/user"
	"github.com/Likhon22/ecom/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()
	dbConfig, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = db.MigrateDB(dbConfig, "./migration")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	middleware := middleware.NewMiddlewares(cnf)

	//repo
	userRepo := repo.NewUserRepo(dbConfig)
	productRepo := repo.NewProductRepo(dbConfig)

	//domains

	userService := user.NewService(userRepo)
	productService := product.NewProductService(productRepo)
	productHandler := productHandler.NewHandler(middleware, productRepo)
	userHandler := userHandler.NewHandler(cnf, userService)

	server := rest.NewServer(productHandler, userHandler, cnf)
	server.StartServer()

}
