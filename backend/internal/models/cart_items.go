package models

type CartItem struct {
	ProductID string  `json:"productId" validate:"required"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity" validate:"required,min=1"`
	Image     string  `json:"image,omitempty"`
	Size      string  `json:"size,omitempty"`
	Subtotal  float64 `json:"subtotal"`
}
