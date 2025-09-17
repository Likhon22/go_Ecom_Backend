package repo

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
	Login(email, password string) (*User, error)
	ListUsers() ([]*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(email string, updatedUser User) (*User, error)
}

type userRepo struct {
	users []*User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}
	return repo
}

func (usr *userRepo) CreateUser(u User) (*User, error) {
	if u.ID != 0 {
		return &u, nil
	}
	u.ID = len(usr.users) + 1
	usr.users = append(usr.users, &u)
	return &u, nil

}
func (usr *userRepo) ListUsers() ([]*User, error) {
	return usr.users, nil
}
func (usr *userRepo) GetUserByEmail(email string) (*User, error) {
	for _, user := range usr.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil

}
func (usr *userRepo) UpdateUser(email string, updatedUser User) (*User, error) {
	for i, user := range usr.users {
		if user.Email == email {
			usr.users[i] = &updatedUser
			return &updatedUser, nil
		}
	}
	return nil, nil
}
func (usr *userRepo) Login(email, password string) (*User, error) {
	for _, user := range usr.users {
		if user.Email == email && user.Password == password {
			return user, nil

		}
	}
	return nil, nil

}
