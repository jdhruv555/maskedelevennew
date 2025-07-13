package utils

import (
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents multiple validation errors
type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword validates password strength
func ValidatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long"
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false, "Password must contain at least one uppercase letter"
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false, "Password must contain at least one lowercase letter"
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false, "Password must contain at least one number"
	}
	return true, ""
}

// ValidatePhone validates phone number format
func ValidatePhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return phoneRegex.MatchString(strings.ReplaceAll(phone, " ", ""))
}

// ValidateRequired validates if a field is required
func ValidateRequired(value, fieldName string) (bool, string) {
	if strings.TrimSpace(value) == "" {
		return false, fieldName + " is required"
	}
	return true, ""
}

// ValidateMinLength validates minimum length
func ValidateMinLength(value string, minLength int, fieldName string) (bool, string) {
	if len(strings.TrimSpace(value)) < minLength {
		return false, fieldName + " must be at least " + string(rune(minLength)) + " characters long"
	}
	return true, ""
}

// ValidateMaxLength validates maximum length
func ValidateMaxLength(value string, maxLength int, fieldName string) (bool, string) {
	if len(value) > maxLength {
		return false, fieldName + " must not exceed " + string(rune(maxLength)) + " characters"
	}
	return true, ""
}

// ValidatePositiveNumber validates if a number is positive
func ValidatePositiveNumber(value float64, fieldName string) (bool, string) {
	if value <= 0 {
		return false, fieldName + " must be a positive number"
	}
	return true, ""
}

// ValidateRange validates if a number is within a range
func ValidateRange(value, min, max float64, fieldName string) (bool, string) {
	if value < min || value > max {
		return false, fieldName + " must be between " + string(rune(int(min))) + " and " + string(rune(int(max)))
	}
	return true, ""
}

// ValidateEnum validates if a value is in a list of allowed values
func ValidateEnum(value string, allowedValues []string, fieldName string) (bool, string) {
	for _, allowed := range allowedValues {
		if value == allowed {
			return true, ""
		}
	}
	return false, fieldName + " must be one of: " + strings.Join(allowedValues, ", ")
}

// ValidatePagination validates pagination parameters
func ValidatePagination(page, limit int) (bool, string) {
	if page < 1 {
		return false, "Page must be greater than 0"
	}
	if limit < 1 || limit > 100 {
		return false, "Limit must be between 1 and 100"
	}
	return true, ""
}

// ValidateUUID validates UUID format
func ValidateUUID(uuid string) bool {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	return uuidRegex.MatchString(strings.ToLower(uuid))
}

// ValidateObjectID validates MongoDB ObjectID format
func ValidateObjectID(id string) bool {
	objectIDRegex := regexp.MustCompile(`^[0-9a-fA-F]{24}$`)
	return objectIDRegex.MatchString(id)
}

// ValidateProductData validates product creation/update data
func ValidateProductData(data map[string]interface{}) []ValidationError {
	var errors []ValidationError

	// Validate required fields
	if name, ok := data["name"].(string); ok {
		if valid, msg := ValidateRequired(name, "name"); !valid {
			errors = append(errors, ValidationError{Field: "name", Message: msg})
		}
		if valid, msg := ValidateMinLength(name, 3, "name"); !valid {
			errors = append(errors, ValidationError{Field: "name", Message: msg})
		}
		if valid, msg := ValidateMaxLength(name, 100, "name"); !valid {
			errors = append(errors, ValidationError{Field: "name", Message: msg})
		}
	} else {
		errors = append(errors, ValidationError{Field: "name", Message: "name is required"})
	}

	if description, ok := data["description"].(string); ok {
		if valid, msg := ValidateMaxLength(description, 1000, "description"); !valid {
			errors = append(errors, ValidationError{Field: "description", Message: msg})
		}
	}

	if price, ok := data["price"].(float64); ok {
		if valid, msg := ValidatePositiveNumber(price, "price"); !valid {
			errors = append(errors, ValidationError{Field: "price", Message: msg})
		}
	} else {
		errors = append(errors, ValidationError{Field: "price", Message: "price is required and must be a number"})
	}

	if category, ok := data["category"].(string); ok {
		if valid, msg := ValidateRequired(category, "category"); !valid {
			errors = append(errors, ValidationError{Field: "category", Message: msg})
		}
	} else {
		errors = append(errors, ValidationError{Field: "category", Message: "category is required"})
	}

	return errors
}

// ValidateUserData validates user registration/update data
func ValidateUserData(data map[string]interface{}) []ValidationError {
	var errors []ValidationError

	if email, ok := data["email"].(string); ok {
		if valid, msg := ValidateRequired(email, "email"); !valid {
			errors = append(errors, ValidationError{Field: "email", Message: msg})
		} else if !ValidateEmail(email) {
			errors = append(errors, ValidationError{Field: "email", Message: "Invalid email format"})
		}
	} else {
		errors = append(errors, ValidationError{Field: "email", Message: "email is required"})
	}

	if password, ok := data["password"].(string); ok {
		if valid, msg := ValidatePassword(password); !valid {
			errors = append(errors, ValidationError{Field: "password", Message: msg})
		}
	} else {
		errors = append(errors, ValidationError{Field: "password", Message: "password is required"})
	}

	if name, ok := data["name"].(string); ok {
		if valid, msg := ValidateRequired(name, "name"); !valid {
			errors = append(errors, ValidationError{Field: "name", Message: msg})
		}
		if valid, msg := ValidateMinLength(name, 2, "name"); !valid {
			errors = append(errors, ValidationError{Field: "name", Message: msg})
		}
		if valid, msg := ValidateMaxLength(name, 50, "name"); !valid {
			errors = append(errors, ValidationError{Field: "name", Message: msg})
		}
	} else {
		errors = append(errors, ValidationError{Field: "name", Message: "name is required"})
	}

	if phone, ok := data["phone"].(string); ok && phone != "" {
		if !ValidatePhone(phone) {
			errors = append(errors, ValidationError{Field: "phone", Message: "Invalid phone number format"})
		}
	}

	return errors
}

// ValidateOrderData validates order creation data
func ValidateOrderData(data map[string]interface{}) []ValidationError {
	var errors []ValidationError

	if items, ok := data["items"].([]interface{}); ok {
		if len(items) == 0 {
			errors = append(errors, ValidationError{Field: "items", Message: "At least one item is required"})
		}
	} else {
		errors = append(errors, ValidationError{Field: "items", Message: "items is required"})
	}

	if shippingAddress, ok := data["shippingAddress"].(string); ok {
		if valid, msg := ValidateRequired(shippingAddress, "shippingAddress"); !valid {
			errors = append(errors, ValidationError{Field: "shippingAddress", Message: msg})
		}
		if valid, msg := ValidateMinLength(shippingAddress, 10, "shippingAddress"); !valid {
			errors = append(errors, ValidationError{Field: "shippingAddress", Message: msg})
		}
	} else {
		errors = append(errors, ValidationError{Field: "shippingAddress", Message: "shippingAddress is required"})
	}

	return errors
}

// SendValidationErrors sends validation errors as a standardized response
func SendValidationErrors(c *fiber.Ctx, errors []ValidationError) error {
	response := ErrorResponse{
		Success:   false,
		Error:     "Validation Error",
		Message:   "Request validation failed",
		Code:      "VALIDATION_ERROR",
		Timestamp: time.Now().UTC(),
		RequestID: c.Get("X-Request-ID"),
	}
	return c.Status(fiber.StatusBadRequest).JSON(response)
} 