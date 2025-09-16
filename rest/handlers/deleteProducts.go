package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Likhon22/ecom/database"
	"github.com/Likhon22/ecom/product"
	"github.com/Likhon22/ecom/utils"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var id product.DeleteRequest
	decoder := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
	err := decoder.Decode(&id)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	isDeleted := database.DeleteProduct(id.ID)
	if !isDeleted {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	productList := database.List()
	utils.SendData(w, productList, http.StatusOK)
}
