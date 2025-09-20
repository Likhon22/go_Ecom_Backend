package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	CreateUser(usr User) (*User, error)
	// Login(email, password string) (*User, error)
	ListUsers() ([]*User, error)
	GetUserByEmail(email string) (*User, error)
	// UpdateUser(email string, updatedUser User) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	repo := &userRepo{
		db: db,
	}
	return repo
}

func (usr *userRepo) CreateUser(u User) (*User, error) {
	query := `
	INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
	VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
	RETURNING id
	`
	rows, err := usr.db.NamedQuery(query, u)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&u.ID)
	}

	return &u, nil

}
func (usr *userRepo) ListUsers() ([]*User, error) {
	var users []*User
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users`
	err := usr.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil

}
func (usr *userRepo) GetUserByEmail(email string) (*User, error) {
	var user User
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1
	LIMIT 1
	`
	err := usr.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err

	}
	return &user, nil
}

// func (usr *userRepo) UpdateUser(email string, updatedUser User) (*User, error) {
// 	for i, user := range usr.users {
// 		if user.Email == email {
// 			usr.users[i] = &updatedUser
// 			return &updatedUser, nil
// 		}
// 	}
// 	return nil, nil
// }
// func (usr *userRepo) Login(email, password string) (*User, error) {
// 	for _, user := range usr.users {
// 		if user.Email == email && user.Password == password {
// 			return user, nil

// 		}
// 	}
// 	return nil, nil
// }
