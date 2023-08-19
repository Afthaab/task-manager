package utility

import (
	"fmt"
	"strings"

	"github.com/afthaab/task-manager/pkg/domain"
	"github.com/go-playground/validator"
)

func ValidateUser(user domain.User) error {
	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))

		for i, validationErr := range validationErrors {
			fieldName := validationErr.Field()
			switch fieldName {
			case "Email":
				errorMessages[i] = "Invalid Email"
				break
			case "Firstname":
				errorMessages[i] = "Invalid Firstname, Minimum 4 letters or Maximum 16 letters required"
				break
			case "Password":
				errorMessages[i] = "Invalid password, Minimum 6 letters or Maximum 16 letters required"
				break
			case "Lastname":
				errorMessages[i] = "Invalid Lastname, Minimum 4 letters or Maximum 16 letters required"
			default:
				errorMessages[i] = "Validation failed"
			}
		}

		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}

	return nil
}
