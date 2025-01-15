package types

import (
	"reflect"
	"strings"
	"time"
	"github.com/google/uuid"
	"github.com/Sahal-P/Go-Auth/utils"
	"github.com/go-playground/validator/v10"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) (*User, error)
}

type User struct {
	ID        uuid.UUID    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// RegisterUserPayload defines the structure for user registration request payload.
type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=20"`
	LastName  string `json:"last_name" validate:"required,min=1,max=20"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

// Validate method calls the validator to validate the struct fields.
func (r *RegisterUserPayload) Validate() map[string]string {
	validate := validator.New() // Create a new validator instance

	// Validate the struct and return customized errors if any
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := validate.Struct(r); err != nil {
		return utils.ValidationErrorMessage(err, r) // Use helper to return friendly error messages
	}
	return nil
}
