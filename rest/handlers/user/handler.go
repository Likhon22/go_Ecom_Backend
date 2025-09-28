package user

import (
	"github.com/Likhon22/ecom/rest/middleware"
)

type Handler struct {
	middleware *middleware.Middlewares
	service    Service
}

func NewHandler(middlewares *middleware.Middlewares, service Service) *Handler {
	return &Handler{
		middleware: middlewares,
		service:    service,
	}
}
