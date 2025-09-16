package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Likhon22/ecom/config"
	"github.com/Likhon22/ecom/rest/middleware"
)

func StartServer(cnf *config.Config) {
	mux := http.NewServeMux()
	mngr := middleware.NewManager()
	mngr.Use(middleware.Logger, middleware.CorsMiddleware)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	initRoutes(mux, mngr)

	wrappedMux := mngr.WrapMux(mux)
	addr := ":" + string(cnf.HttpPort)

	fmt.Println("Server started on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {

		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
