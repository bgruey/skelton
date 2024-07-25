package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID          int64  `gorm:"unique;primaryKey;autoIncrement" json:"-"`
	UUID        string `gorm:"unique_index"`
	CreatedAt   int64  `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli" json:"-"`
	ValidatedAt *int64 `json:"-"`

	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique_index"`
	PasswordHash string `json:"-"`
}

func NewUser(name string, email string) *User {
	ret := new(User)

	ret.Name = name
	ret.Email = email

	ret.Update()

	return ret
}

func (u *User) Update() {
	u.UUID = uuid.NewString()
}
