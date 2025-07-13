package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/Shrey-Yash/Masked11/internal/services"
	"github.com/Shrey-Yash/Masked11/internal/utils"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	key, err := utils.GetCartKey(c)
	if err != nil {
		return err
	}

	erro := h.OrderService.CreateOrder(key)
	if erro != nil {
		return fiber.NewError(fiber.StatusInternalServerError, erro.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Order Created Successfully",
	})
}

func (h *OrderHandler) GetOrdersByUserID(c *fiber.Ctx) error {
	uid, ok := utils.GetCartKey(c)
	if ok != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	orders, err := h.OrderService.GetOrdersByUserID(context.Background(), uid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(orders)
}

func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("id")
	userID, _ := utils.GetCartKey(c)
	isAdmin, _ := c.Locals("isAdmin").(bool)

	order, err := h.OrderService.GetOrderByID(context.Background(), orderID, isAdmin, userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if !isAdmin && order.UserID != userID {
		return fiber.NewError(fiber.StatusForbidden, "Unauthorized attempt to access this order.")
	}
	return c.JSON(order)
}

func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	orderID := c.Params("id")
	type payload struct {
		Status string `json:"status"`
	}

	var p payload
	if err := c.BodyParser(&p); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Body.")
	}

	if err := h.OrderService.UpdateOrderStatus(orderID, p.Status); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "OrderStatusUpdated.",
	})
}

func (h *OrderHandler) CancelOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	userID, err := utils.GetCartKey(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not Authenticated")
	}

	if err := h.OrderService.CancelOrder(c.Context(), userID, orderID); err != nil {
		if err.Error() == "unauthorized access to cancel order" {
			return fiber.NewError(fiber.StatusForbidden, "You are not allowed to cancel this order.")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to cancel this order.")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order cancelled successfully.",
	})
}

func (h *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	if err := h.OrderService.DeleteOrder(orderID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Order Deleted",
	})
}
