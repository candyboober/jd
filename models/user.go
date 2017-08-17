package models

import (
	"golang.org/x/crypto/bcrypt"
)

var PasswordSalt = []byte("passwordsalt")

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"size:20; not null"`
	Password []byte `json:"-" gorm:"size:255; not null"`
}

func (this *User) HashingPassword() {
	var err error
	this.Password, err = bcrypt.GenerateFromPassword(this.Password, 10)
	if err != nil {
		panic(err)
	}
}
