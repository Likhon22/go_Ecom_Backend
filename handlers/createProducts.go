package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Likhon22/ecom/product"
	"github.com/Likhon22/ecom/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)
	utils.SendData(w, newProduct, http.StatusCreated)

}
