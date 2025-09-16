package handlers

import (
	"net/http"
	"strconv"

	"github.com/Likhon22/ecom/database"
	"github.com/Likhon22/ecom/utils"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	searchedProduct := database.GetProductByID(id)

	utils.SendData(w, searchedProduct, http.StatusOK)

}
