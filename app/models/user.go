package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	CustomModel

	Username string
	Email    string
	Password []byte `json:",omitempty"`

	Blogs []Blog `json:",omitempty"`
}

// type SafeUser struct {
// 	CustomModel

// 	Username string
// 	Email    string
// }

func (u *User) UserToSafeUser() User {
	return User{
		Model:       u.Model,
		CustomModel: u.CustomModel,
		Username:    u.Username,
		Email:       u.Email,
	}
}
