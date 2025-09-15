package cmd

import (
	"fmt"
	"net/http"

	"github.com/Likhon22/ecom/handlers"
	"github.com/Likhon22/ecom/middleware"
)

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.Handle("GET /getProducts", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /createProduct", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("PATCH /updateProduct", http.HandlerFunc(handlers.UpdateProduct))
	mux.Handle("DELETE /deleteProduct", http.HandlerFunc(handlers.DeleteProduct))
	handler := middleware.CorsMiddleware(mux)
	fmt.Println("Server started on port 3000")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {

		fmt.Println("Error starting server:", err)
	}
}
