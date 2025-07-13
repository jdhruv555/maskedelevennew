
package interfaces

import "github.com/Shrey-Yash/Masked11/internal/models"

type CartRepository interface {
	GetCart(key string) (*models.Cart, error)
	SetCart(key string, cart *models.Cart) error
	DeleteCart(key string) error
}