package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int `json:"id"`
	Title       string
	Description string
	Price       float64
	Image       string
}

var productList []Product

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)

}

func getProducts(w http.ResponseWriter, r *http.Request) {

	sendData(w, productList, http.StatusOK)

}

func createProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendData(w, newProduct, http.StatusCreated)

}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	for _, product := range productList {
		if product.ID == updatedProduct.ID {
			product.Title = updatedProduct.Title
			product.Description = updatedProduct.Description
			product.Price = updatedProduct.Price
			product.Image = updatedProduct.Image
			sendData(w, product, http.StatusOK)
			return
		}

	}

}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	var id int
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	for _, product := range productList {
		if product.ID == id {
			productList = append(productList[:id-1], productList[id:]...)
			sendData(w, productList, http.StatusOK)
			return
		}

	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.Handle("GET /getProducts", http.HandlerFunc(getProducts))
	mux.Handle("POST /createProduct", http.HandlerFunc(createProduct))
	mux.Handle("PATCH /updateProduct", http.HandlerFunc(updateProduct))
	mux.Handle("DELETE /deleteProduct", http.HandlerFunc(deleteProduct))
	handler := corsMiddleware(mux)
	fmt.Println("Server started on port 3000")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {

		fmt.Println("Error starting server:", err)
	}

}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "Description 1",
		Price:       100,
		Image:       "Image 1",
	}
	product2 := Product{
		ID:          2,
		Title:       "Product 2",
		Description: "Description 2",
		Price:       200,
		Image:       "Image 2",
	}
	product3 := Product{
		ID:          3,
		Title:       "Product 3",
		Description: "Description 3",
		Price:       300,
		Image:       "Image 3",
	}
	productList = append(productList, product1, product2, product3)

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
