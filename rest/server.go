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

	cnf *config.Config
}

func NewServer(productHandler *product.Handler, userHandler *user.Handler, cnf *config.Config) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,

		cnf: cnf,
	}
}

func (server *Server) StartServer() {
	mux := http.NewServeMux()
	mngr := middleware.NewManager()
	mngr.Use(middleware.Logger, middleware.CorsMiddleware)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	server.productHandler.ProductRoutes(mux, mngr)
	server.userHandler.UserRoutes(mux, mngr)

	wrappedMux := mngr.WrapMux(mux)
	addr := ":" + string(server.cnf.HttpPort)

	fmt.Println("Server started on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {

		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
