package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Likhon22/ecom/database"
	"github.com/Likhon22/ecom/product"
	"github.com/Likhon22/ecom/utils"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	header := r.Header.Get("Authorization")
	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return

	}
	headerArr := strings.Split(header, " ")
	if len(headerArr) != 2 || headerArr[0] != "Bearer" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	accessToken := headerArr[1]
	log.Println(accessToken)

	var newProduct product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	database.StoreProduct(newProduct)
	utils.SendData(w, newProduct, http.StatusCreated)

}
