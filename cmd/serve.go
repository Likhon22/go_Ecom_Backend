package cmd

import (
	"fmt"
	"net/http"

	"github.com/Likhon22/ecom/handlers"
	"github.com/Likhon22/ecom/middleware"
)

func Serve() {
	mux := http.NewServeMux()
	mngr := middleware.NewManager()
	mngr.Use(middleware.CorsMiddleware, middleware.Logger)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.Handle("GET /products", mngr.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /product/{id}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("POST /product", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("PATCH /product", http.HandlerFunc(handlers.UpdateProduct))
	mux.Handle("DELETE /product", http.HandlerFunc(handlers.DeleteProduct))

	fmt.Println("Server started on port 3000")
	handler := middleware.CorsMiddleware(mux)
	err := http.ListenAndServe(":3000", handler)
	if err != nil {

		fmt.Println("Error starting server:", err)
	}
}
