package cmd

import (
	"fmt"
	"net/http"

	"github.com/Likhon22/ecom/middleware"
)

func Serve() {
	mux := http.NewServeMux()
	mngr := middleware.NewManager()
	mngr.Use(middleware.Logger, middleware.CorsMiddleware)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	initRoutes(mux, mngr)

	fmt.Println("Server started on port 3000")


	
	wrappedMux := mngr.WrapMux(mux)
	err := http.ListenAndServe(":3000", wrappedMux)
	if err != nil {

		fmt.Println("Error starting server:", err)
	}
}
