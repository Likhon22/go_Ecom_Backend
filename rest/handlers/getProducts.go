package handlers

import (
	"net/http"

	"github.com/Likhon22/ecom/database"
	"github.com/Likhon22/ecom/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	utils.SendData(w, database.List(), http.StatusOK)

}
