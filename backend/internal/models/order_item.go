package models

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"orderId"`
	ProductID string    `json:"productId"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	Size      string    `json:"size,omitempty"`
	Image     string    `json:"image,omitempty"`
	Subtotal  float64   `json:"subtotal"`
}
