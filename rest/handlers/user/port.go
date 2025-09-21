package user

import "github.com/Likhon22/ecom/domain"

type Service interface {
	CreateUser(usr domain.User) (*domain.User, error)
	Login(email, password string) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	UpdateUser(email string, updatedUser domain.User) (*domain.User, error)
}
