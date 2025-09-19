package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Likhon22/ecom/repo"
	"github.com/Likhon22/ecom/utils"
)

type ReqCreateProduct struct {
	Title       string
	Description string
	Price       float64
	Image       string
}

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

	var newProduct ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	created, err := h.productRepo.Create(repo.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		Image:       newProduct.Image,
	})
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, created, http.StatusCreated)

}
