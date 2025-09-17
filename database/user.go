package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var users []User

func (u User) StoreUser() User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(users) + 1
	users = append(users, u)
	return u
}

func ListUsers() []User {
	return users
}
func GetUserByEmail(email string) *User {
	for _, user := range users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}
func UpdateUser(email string, updatedUser User) *User {
	for i, user := range users {
		if user.Email == email {
			users[i] = updatedUser
			return &updatedUser
		}
	}
	return nil
}

func Login(email, password string) *User {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return &user

		}
	}
	return nil
}
