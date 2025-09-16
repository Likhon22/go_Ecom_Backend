package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Likhon22/ecom/database"
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
	updatedStatus := database.UpdateProduct(updatedProduct)
	if !updatedStatus {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, updatedProduct, http.StatusOK)

}
