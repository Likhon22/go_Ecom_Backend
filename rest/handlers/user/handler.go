package user

import (
	"github.com/Likhon22/ecom/config"
)

type Handler struct {
	cnf     *config.Config
	service Service
}

func NewHandler(cnf *config.Config, service Service) *Handler {
	return &Handler{
		cnf:     cnf,
		service: service,
	}

}
