package product

import (
	"net/http"

	"github.com/Likhon22/ecom/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.GetAll()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, products, http.StatusOK)

}
