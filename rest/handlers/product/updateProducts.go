package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Likhon22/ecom/repo"
	"github.com/Likhon22/ecom/utils"
)

type ReqUpdateProduct struct {
	ID          int
	Title       string
	Description string
	Price       float64
	Image       string
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct *ReqUpdateProduct

	pId, _ := strconv.Atoi(r.PathValue("id"))
	updatedProduct.ID = pId
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	created, err := h.productRepo.Update(repo.Product{
		ID:          updatedProduct.ID,
		Title:       updatedProduct.Title,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		Image:       updatedProduct.Image,
	})
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, created, http.StatusOK)

}
