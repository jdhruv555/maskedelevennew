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
	GetAllProductsWithFilters(ctx context.Context, filters map[string]interface{}) ([]*models.Product, int64, error)
	SearchProducts(ctx context.Context, filters map[string]interface{}) ([]*models.Product, int64, error)
	GetProductsByCategory(ctx context.Context, filters map[string]interface{}) ([]*models.Product, int64, error)
	GetFeaturedProducts(ctx context.Context, limit int) ([]*models.Product, error)
	GetProductCategories(ctx context.Context) ([]string, error)
	GetProductsByPriceRange(ctx context.Context, filters map[string]interface{}) ([]*models.Product, int64, error)
	GetProductsInStock(ctx context.Context, filters map[string]interface{}) ([]*models.Product, int64, error)
	GetNewArrivals(ctx context.Context, limit int) ([]*models.Product, error)
	GetBestSellers(ctx context.Context, limit int) ([]*models.Product, error)
	GetRelatedProducts(ctx context.Context, productID string, limit int) ([]*models.Product, error)
}
