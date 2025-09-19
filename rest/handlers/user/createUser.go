package user

import (
	"encoding/json"
	"net/http"

	"github.com/Likhon22/ecom/repo"
	"github.com/Likhon22/ecom/utils"
)

type CreateUserRequest struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser CreateUserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	created, err := h.UserRepo.CreateUser(repo.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, created, http.StatusCreated)

}
