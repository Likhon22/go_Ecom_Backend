package middleware

import "github.com/Likhon22/ecom/config"

type Middlewares struct {
	config *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{config: cnf}
}
