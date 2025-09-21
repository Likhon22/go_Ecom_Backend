package user

import (
	"database/sql"
	"fmt"

	"github.com/Likhon22/ecom/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (usr *service) CreateUser(u domain.User) (*domain.User, error) {
	user, err := usr.usrRepo.CreateUser(u)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}
	return user, nil

}
func (usr *service) ListUsers() ([]*domain.User, error) {
	user, err := usr.usrRepo.ListUsers()
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (usr *service) GetUserByEmail(email string) (*domain.User, error) {
	user, err := usr.usrRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (usr *service) UpdateUser(email string, updatedUser domain.User) (*domain.User, error) {
	user, err := usr.usrRepo.UpdateUser(email, updatedUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (usr *service) Login(email, password string) (*domain.User, error) {
	user, err := usr.usrRepo.Login(email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid email or password")
		}
		return nil, err
	}
	return user, nil
}
