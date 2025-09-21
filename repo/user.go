package repo

import (
	"database/sql"
	"fmt"

	"github.com/Likhon22/ecom/domain"
	"github.com/Likhon22/ecom/user"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}
type UserRepo interface {
	user.UserRepo
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	repo := &userRepo{
		db: db,
	}
	return repo
}

func (usr *userRepo) CreateUser(u domain.User) (*domain.User, error) {
	query := `
	INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
	VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
	RETURNING id
	`
	rows, err := usr.db.NamedQuery(query, u)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&u.ID)
	}

	return &u, nil

}
func (usr *userRepo) ListUsers() ([]*domain.User, error) {
	var users []*domain.User
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users`
	err := usr.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil

}
func (usr *userRepo) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
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

func (usr *userRepo) UpdateUser(email string, updatedUser domain.User) (*domain.User, error) {
	query := `
		UPDATE users
		SET first_name=$1, last_name=$2, updated_at=NOW()
		WHERE email=$3
	`

	result, err := usr.db.Exec(query, updatedUser.FirstName, updatedUser.LastName, email)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		// No user found with this email
		return nil, nil
	}

	return &updatedUser, nil
}
func (usr *userRepo) Login(email, password string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1
		LIMIT 1`
	err := usr.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err

	}
	if user.Password == password {
		return &user, nil
	}
	return nil, nil
}
