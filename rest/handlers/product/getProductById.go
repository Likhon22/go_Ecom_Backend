package product

import (
	"net/http"
	"strconv"

	"github.com/Likhon22/ecom/utils"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	searchedProduct, err := h.productRepo.GetByID(id)
	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		return
	}
	if searchedProduct == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return

	}

	utils.SendData(w, searchedProduct, http.StatusOK)

}
