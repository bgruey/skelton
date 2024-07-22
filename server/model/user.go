package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID          int64 `gorm:"unique;primaryKey;autoIncrement" json:"-"`
	UUID        string
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UpdatedAt   int64 `gorm:"autoUpdateTime:milli"`
	ValidatedAt *int64

	Name     string `json:"name"`
	Email    string `json:"email" gorm:"index"`
	Password string
}

func NewUser(name string, email string) *User {
	ret := new(User)

	ret.Name = name
	ret.Email = email

	ret.UUID = uuid.NewString()

	return ret
}

func (u *User) Update() {
	u.UUID = uuid.NewString()
}
