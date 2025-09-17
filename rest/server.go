package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Likhon22/ecom/config"
	"github.com/Likhon22/ecom/rest/handlers/product"
	"github.com/Likhon22/ecom/rest/handlers/user"
	"github.com/Likhon22/ecom/rest/middleware"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		productHandler: product.NewHandler(),
		userHandler:    user.NewHandler(),
	}
}

func (server *Server) StartServer(cnf *config.Config) {
	mux := http.NewServeMux()
	mngr := middleware.NewManager()
	mngr.Use(middleware.Logger, middleware.CorsMiddleware)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	server.productHandler.ProductRoutes(mux, mngr)
	server.userHandler.UserRoutes(mux, mngr)

	wrappedMux := mngr.WrapMux(mux)
	addr := ":" + string(cnf.HttpPort)

	fmt.Println("Server started on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {

		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
