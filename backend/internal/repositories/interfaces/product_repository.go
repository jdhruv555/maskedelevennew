package interfaces

import (
	"context"

	"github.com/Shrey-Yash/Masked11/internal/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	UpdateProduct(ctx context.Context, id string, updates map[string]interface{}) error
	DeleteProduct(ctx context.Context, id string) error
}
