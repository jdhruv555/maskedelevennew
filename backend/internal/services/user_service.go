package services

import (
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/Shrey-Yash/Masked11/internal/models"
)

var validate = validator.New()

func ValidateUserInput(u *models.User) map[string]string {
	err := validate.Struct(u)
	if err == nil {
		return nil
	}

	errs := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())

		switch e.Tag() {
		case "required":
			errs[field] = field + " is required"
		case "email":
			errs[field] = "Enter a valid E-Mail address."
		case "min":
			if field == "password" {
				errs[field] = "Password must have atleast 6 characters."
			} else {
				errs[field] = field + " is too short."
			}
		default:
			errs[field] = "Invalid " + field
		}
	}
	return errs
}
