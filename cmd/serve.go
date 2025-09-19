package cmd

import (
	"github.com/Likhon22/ecom/config"
	"github.com/Likhon22/ecom/repo"
	"github.com/Likhon22/ecom/rest"
	"github.com/Likhon22/ecom/rest/handlers/product"
	"github.com/Likhon22/ecom/rest/handlers/review"
	"github.com/Likhon22/ecom/rest/handlers/user"
	"github.com/Likhon22/ecom/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()
	middleware := middleware.NewMiddlewares(cnf)
	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()
	productHandler := product.NewHandler(middleware, productRepo)
	userHandler := user.NewHandler(userRepo)
	reviewHandler := review.NewHandler()
	server := rest.NewServer(productHandler, userHandler, reviewHandler, cnf)
	server.StartServer()

}
