package interfaces

import (
	"context"

	"github.com/Shrey-Yash/Masked11/internal/models"
)

type OrderRepository interface {
	CreateOrder(order *models.Order, items []models.OrderItem) error
	GetOrdersByUserID(ctx context.Context, userID string) ([]models.Order, error)
	GetOrderByID(ctx context.Context, orderID string) (*models.Order, error)
	UpdateOrderStatus(orderID string, status string) error
	DeleteOrder(orderID string) error
	CancelOrder(orderID string) error
}