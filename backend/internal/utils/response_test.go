package utils

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSuccessResponse(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c *fiber.Ctx) error {
		return SuccessResponse(c, fiber.StatusOK, "Test successful", fiber.Map{"key": "value"})
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var response Response
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	assert.True(t, response.Success)
	assert.Equal(t, "Test successful", response.Message)
	assert.NotNil(t, response.Data)
	assert.NotZero(t, response.Timestamp)
}

func TestSendErrorResponse(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c *fiber.Ctx) error {
		return SendErrorResponse(c, fiber.StatusBadRequest, "Validation Error", "Invalid input")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var response ErrorResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	assert.False(t, response.Success)
	assert.Equal(t, "Validation Error", response.Error)
	assert.Equal(t, "Invalid input", response.Message)
	assert.NotZero(t, response.Timestamp)
}

func TestValidationErrorResponse(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c *fiber.Ctx) error {
		errors := []ValidationError{
			{Field: "email", Message: "Invalid email format"},
			{Field: "password", Message: "Password too short"},
		}
		return SendValidationErrors(c, errors)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var response ErrorResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	assert.False(t, response.Success)
	assert.Equal(t, "Validation Error", response.Error)
	assert.Equal(t, "Request validation failed", response.Message)
	assert.Equal(t, "VALIDATION_ERROR", response.Code)
}

func TestServerErrorResponse(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c *fiber.Ctx) error {
		return ServerErrorResponse(c, "Database connection failed")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	var response ErrorResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	assert.False(t, response.Success)
	assert.Equal(t, "Internal Server Error", response.Error)
	assert.Equal(t, "Database connection failed", response.Message)
}

func TestNotFoundResponse(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c *fiber.Ctx) error {
		return NotFoundResponse(c, "Resource not found")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	var response ErrorResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	assert.False(t, response.Success)
	assert.Equal(t, "Not Found", response.Error)
	assert.Equal(t, "Resource not found", response.Message)
}

func TestUnauthorizedResponse(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c *fiber.Ctx) error {
		return UnauthorizedResponse(c, "Invalid credentials")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	var response ErrorResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	assert.False(t, response.Success)
	assert.Equal(t, "Unauthorized", response.Error)
	assert.Equal(t, "Invalid credentials", response.Message)
}

func TestForbiddenResponse(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c *fiber.Ctx) error {
		return ForbiddenResponse(c, "Insufficient permissions")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)

	var response ErrorResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	assert.False(t, response.Success)
	assert.Equal(t, "Forbidden", response.Error)
	assert.Equal(t, "Insufficient permissions", response.Message)
}

func TestResponseWithRequestID(t *testing.T) {
	app := fiber.New()
	
	app.Get("/test", func(c *fiber.Ctx) error {
		return SuccessResponse(c, fiber.StatusOK, "Test successful", nil)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-Request-ID", "test-request-id")
	resp, _ := app.Test(req)

	body, _ := io.ReadAll(resp.Body)
	var response Response
	json.Unmarshal(body, &response)

	assert.Equal(t, "test-request-id", response.RequestID)
} 