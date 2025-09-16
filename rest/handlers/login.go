package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Likhon22/ecom/database"
	"github.com/Likhon22/ecom/utils"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLoginUser ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLoginUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	user := database.Login(reqLoginUser.Email, reqLoginUser.Password)
	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	utils.SendData(w, user, http.StatusOK)

}
