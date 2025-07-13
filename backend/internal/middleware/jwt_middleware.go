package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/Shrey-Yash/Masked11/internal/utils"
	"github.com/Shrey-Yash/Masked11/internal/database"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("token")
		if cookie == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing auth token")
		}

		claims, err := utils.ParseJWT(cookie)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		jti, ok := claims["jti"].(string)
		if !ok || strings.TrimSpace(jti) == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid session token")
		}

		exists, err := database.Redis.Exists(database.Ctx, jti).Result()
		if err != nil || exists == 0 {
			return fiber.NewError(fiber.StatusUnauthorized, "Session expired or invalid")
		}

		c.Locals("userID", claims["sub"])
		c.Locals("userEmail", claims["email"])

		return c.Next()
	}
}
