package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID   `json:"id"`
	UserID    string      `json:"userId"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	Items     []OrderItem `json:"items"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}
