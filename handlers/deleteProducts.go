package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	index := -1
	for i, product := range product.ProductList {
		if product.ID == id.ID {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	product.ProductList = append(product.ProductList[:index], product.ProductList[index+1:]...)
	utils.SendData(w, product.ProductList, http.StatusOK)
}
