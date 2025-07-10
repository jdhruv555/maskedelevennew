package handlers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/services"
)

type UserHandler struct {
	UserService *services.AuthService
}

func NewUserHandler(service *services.AuthService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	user, err := h.UserService.GetUserByID(userId.(string))
	if err != nil {
		log.Println("GetUser error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch user")
	}

	user.Password = ""
	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	var updates map[string]interface{}
	if err := json.Unmarshal(c.Body(), &updates); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	updatedUser := &models.User{}
	if name, ok := updates["name"].(string); ok {
		updatedUser.Name = strings.TrimSpace(name)
	}
	if phone, ok := updates["phone"].(string); ok {
		updatedUser.Phone = strings.TrimSpace(phone)
	}
	if address, ok := updates["address"].(string); ok {
		updatedUser.Address = strings.TrimSpace(address)
	}

	err := h.UserService.UpdateUser(userId.(string), updatedUser)
	if err != nil {
		log.Println("UpdateUser error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	err := h.UserService.DeleteUser(userId.(string))
	if err != nil {
		log.Println("DeleteUser error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete user")
	}

	c.ClearCookie("token")
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
