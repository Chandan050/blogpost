package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       uint   `json:"id"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Phone    string `json:"phone"`
}

func (user *User) SetPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// Handle error (e.g., log it, return an error, etc.)
		return
	}

	user.Password = hash
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
