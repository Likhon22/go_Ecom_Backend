package user

import "github.com/Likhon22/ecom/repo"

type Handler struct {
	UserRepo repo.UserRepo
}

func NewHandler(repo repo.UserRepo) *Handler {
	return &Handler{
		UserRepo: repo,
	}

}
