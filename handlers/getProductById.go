package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Likhon22/ecom/product"
	"github.com/Likhon22/ecom/utils"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var searchedProduct product.Product
	fmt.Println(product.ProductList)
	for _, v := range product.ProductList {
		if v.ID == id {
			searchedProduct = v

		}

	}
	if searchedProduct.ID == 0 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, searchedProduct, http.StatusOK)

}
