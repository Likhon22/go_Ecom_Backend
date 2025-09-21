package user

import (
	"encoding/json"
	"net/http"

	"github.com/Likhon22/ecom/config"
	"github.com/Likhon22/ecom/domain"
	"github.com/Likhon22/ecom/utils"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLoginUser domain.ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLoginUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	user, err := h.service.Login(reqLoginUser.Email, reqLoginUser.Password)
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return

	}
	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	accessToken, err := utils.CreateJwt(config.GetConfig().SecretKey, utils.Payload{
		Sub:         user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusOK)

}
