package dto

import (
	// validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Usertype string

const (
	// user roles
	JobSeeker Usertype = "Employee"
	Employer  Usertype = "Recuriter"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	RoleId   int    `json:"role_id"`
}

func (u User) Validate() error {

	err := validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, is.Alphanumeric),
		validation.Field(&u.Phone, validation.Required, validation.Length(0, 10)),
		validation.Field(&u.RoleId, validation.Required),
	)

	return err

}
