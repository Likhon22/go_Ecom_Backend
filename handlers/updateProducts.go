package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Likhon22/ecom/product"
	"github.com/Likhon22/ecom/utils"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	for _, product := range product.ProductList {
		if product.ID == updatedProduct.ID {
			product.Title = updatedProduct.Title
			product.Description = updatedProduct.Description
			product.Price = updatedProduct.Price
			product.Image = updatedProduct.Image
			utils.SendData(w, product, http.StatusOK)
			return
		}

	}

}
