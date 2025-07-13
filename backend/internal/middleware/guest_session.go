package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GuestSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Cookies("guest_sid")

		if sessionID == "" {
			sessionID = uuid.New().String()
			c.Cookie(&fiber.Cookie{
				Name:     "guest_sid",
				Value:    sessionID,
				Path:     "/",
				HTTPOnly: true,
				Secure:   true,
				SameSite: "Lax",
			})
		}

		c.Locals("sessionID", sessionID)
		return c.Next()
	}
}
