package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetCartKey(c *fiber.Ctx) (string, error) {
	if uid, ok := c.Locals("userID").(string); ok && uid != "" {
		return fmt.Sprintf("user:%s:cart", uid), nil
	}
	if sid := c.Locals("sessionID"); sid != "" {
		return fmt.Sprintf("guest:%s:cart", sid), nil
	}
	return "", fiber.NewError(fiber.StatusUnauthorized, "Unable to identify session")
}
