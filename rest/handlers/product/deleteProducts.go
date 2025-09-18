package product

import (
	"net/http"
	"strconv"

	"github.com/Likhon22/ecom/utils"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
	}

	isDeleted, err := h.productRepo.Delete(pId)
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}
	if !isDeleted {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	productList, err := h.productRepo.GetAll()
	if err != nil {
		http.Error(w, "Error fetching product list", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, productList, http.StatusOK)
}
