package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Shrey-Yash/Masked11/internal/utils"
)

func AdminOnly() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Cookies("token")
        if token == "" {
            return fiber.NewError(fiber.StatusUnauthorized, "Missing auth token")
        }

        claims, err := utils.ParseJWT(token)
        if err != nil {
            return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
        }

        role, ok := claims["role"].(string)
        if !ok || role != "admin" {
            return fiber.NewError(fiber.StatusForbidden, "Admin access required")
        }

        return c.Next()
    }
}