package product

import (
	"net/http"
	"strconv"

	"github.com/Likhon22/ecom/database"
	"github.com/Likhon22/ecom/utils"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
	}

	isDeleted := database.DeleteProduct(pId)
	if !isDeleted {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	productList := database.List()
	utils.SendData(w, productList, http.StatusOK)
}
