package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Likhon22/ecom/database"
	"github.com/Likhon22/ecom/product"
	"github.com/Likhon22/ecom/utils"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updatedProduct product.Product

	updatedProduct.ID, _ = strconv.Atoi(r.PathValue("id"))
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	updatedProduct = database.UpdateProduct(updatedProduct)

	utils.SendData(w, updatedProduct, http.StatusOK)

}
