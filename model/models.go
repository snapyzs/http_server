package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Users struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Lists struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (u *Users) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.Password == "")), validation.Length(6, 100)),
	)
}
