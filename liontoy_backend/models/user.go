package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"Password"`
	Quyen    string `json:"quyen"`
}

func (user *User) SetPassword(password string) {
	hashedpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedpassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
