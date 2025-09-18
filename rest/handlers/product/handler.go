package product

import (
	"github.com/Likhon22/ecom/repo"
	"github.com/Likhon22/ecom/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(middlewares *middleware.Middlewares, repo repo.ProductRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: repo,
	}
}
