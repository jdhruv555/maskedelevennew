package handlers

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/repositories/redis"
	"github.com/Shrey-Yash/Masked11/internal/services"
	"github.com/Shrey-Yash/Masked11/internal/utils"
)

type AuthHandler struct {
	AuthService *services.AuthService
	SessionRepo *redisrepo.SessionRepository
}

func NewAuthHandler(authService *services.AuthService, sessionRepo *redisrepo.SessionRepository) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		SessionRepo: sessionRepo,
	}
}

func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	var user models.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		log.Println("Invalid request body:", err)
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid Request Format", "Invalid request format.")
	}
	
	// Validate user data
	userData := map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
		"name":     user.Name,
	}
	
	if errors := utils.ValidateUserData(userData); len(errors) > 0 {
		return utils.SendValidationErrors(c, errors)
	}
	
	user.Role = "user"

	if err := h.AuthService.RegisterUser(&user); err != nil {
		log.Println("Registration failed:", err)
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Registration Failed", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", fiber.Map{
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
	})
}

func (h *AuthHandler) LoginUser(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.AuthService.LoginUser(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, jti, err := utils.GenerateJWT(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	err = h.SessionRepo.StoreSession(jti, user.ID.Hex(), 7*24*time.Hour)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Session storage failed",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   os.Getenv("APP_ENV") == "production",
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	if cookie == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "No auth token provided")
	}

	claims, err := utils.ParseJWT(cookie)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}
	jti := claims["jti"].(string)

	if err := h.SessionRepo.DeleteSession(jti); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete session")
	}

	c.ClearCookie("token")

	return c.JSON(fiber.Map{"message": "Logged out successfully"})
}
