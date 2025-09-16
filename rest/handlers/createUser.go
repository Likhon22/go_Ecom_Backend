package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Likhon22/ecom/database"
	"github.com/Likhon22/ecom/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	newUser.StoreUser()
	utils.SendData(w, newUser, http.StatusCreated)

}
