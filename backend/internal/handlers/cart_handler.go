package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/repositories/interfaces"
	"github.com/Shrey-Yash/Masked11/internal/utils"
)

type CartHandler struct {
	CartRepo interfaces.CartRepository
}

func NewCartHandler(cartRepo interfaces.CartRepository) *CartHandler {
	return &CartHandler{CartRepo: cartRepo}
}

func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	key, err := utils.GetCartKey(c)
	if err != nil {
		return err
	}

	var item models.CartItem
	if err := c.BodyParser(&item); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	cart, err := h.CartRepo.GetCart(key)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch cart")
	}
	if cart == nil {
		cart = &models.Cart{
			UserID:    key,
			Items:     []models.CartItem{},
			UpdatedAt: time.Now(),
		}
	}

	updated := false
	for i, ci := range cart.Items {
		if ci.ProductID == item.ProductID && ci.Size == item.Size {
			cart.Items[i].Quantity += item.Quantity
			cart.Items[i].Subtotal = float64(cart.Items[i].Quantity) * cart.Items[i].Price
			updated = true
			break
		}
	}

	if !updated {
		item.Subtotal = float64(item.Quantity) * item.Price
		cart.Items = append(cart.Items, item)
	}

	cart.UpdatedAt = time.Now()
	if err := h.CartRepo.SetCart(key, cart); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update cart")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item added to cart"})
}

func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	key, err := utils.GetCartKey(c)
	if err != nil {
		return err
	}

	cart, err := h.CartRepo.GetCart(key)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch cart")
	}
	if cart == nil {
		return c.JSON(&models.Cart{UserID: key, Items: []models.CartItem{}})
	}

	return c.JSON(cart)
}

func (h *CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	key, err := utils.GetCartKey(c)
	if err != nil {
		return err
	}

	productID := c.Params("id")
	size := c.Query("size")

	cart, err := h.CartRepo.GetCart(key)
	if err != nil || cart == nil {
		return fiber.NewError(fiber.StatusNotFound, "Cart not found")
	}

	filtered := []models.CartItem{}
	for _, item := range cart.Items {
		if item.ProductID != productID || (size != "" && item.Size != size) {
			filtered = append(filtered, item)
		}
	}
	cart.Items = filtered

	if err := h.CartRepo.SetCart(key, cart); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update cart")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item removed from cart"})
}

func (h *CartHandler) ClearCart(c *fiber.Ctx) error {
	key, err := utils.GetCartKey(c)
	if err != nil {
		return err
	}

	if err := h.CartRepo.DeleteCart(key); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to clear cart")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Cart cleared"})
}
