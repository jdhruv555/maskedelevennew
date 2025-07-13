package models

import "time"

type Cart struct {
	UserID    string        `json:"userId,omitempty"`
	Items     []CartItem    `json:"items"`
	UpdatedAt time.Time     `json:"updatedAt"`
}
