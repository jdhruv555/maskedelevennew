package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Response represents a standardized API response structure
type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	RequestID string      `json:"request_id,omitempty"`
}

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Success   bool      `json:"success"`
	Error     string    `json:"error"`
	Message   string    `json:"message,omitempty"`
	Code      string    `json:"code,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	RequestID string    `json:"request_id,omitempty"`
}

// SuccessResponse sends a standardized success response
func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := Response{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UTC(),
		RequestID: c.Get("X-Request-ID"),
	}
	return c.Status(statusCode).JSON(response)
}

// SendErrorResponse sends a standardized error response
func SendErrorResponse(c *fiber.Ctx, statusCode int, errorMsg string, message string) error {
	response := ErrorResponse{
		Success:   false,
		Error:     errorMsg,
		Message:   message,
		Timestamp: time.Now().UTC(),
		RequestID: c.Get("X-Request-ID"),
	}
	return c.Status(statusCode).JSON(response)
}

// ValidationErrorResponse sends a standardized validation error response
func ValidationErrorResponse(c *fiber.Ctx, errors map[string]string) error {
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

// ServerErrorResponse sends a standardized server error response
func ServerErrorResponse(c *fiber.Ctx, message string) error {
	return SendErrorResponse(c, fiber.StatusInternalServerError, "Internal Server Error", message)
}

// NotFoundResponse sends a standardized not found response
func NotFoundResponse(c *fiber.Ctx, message string) error {
	return SendErrorResponse(c, fiber.StatusNotFound, "Not Found", message)
}

// UnauthorizedResponse sends a standardized unauthorized response
func UnauthorizedResponse(c *fiber.Ctx, message string) error {
	return SendErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized", message)
}

// ForbiddenResponse sends a standardized forbidden response
func ForbiddenResponse(c *fiber.Ctx, message string) error {
	return SendErrorResponse(c, fiber.StatusForbidden, "Forbidden", message)
} 