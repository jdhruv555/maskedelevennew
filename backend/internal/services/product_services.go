package services

import (
	"context"
	"strings"
	"time"

	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/repositories/interfaces"
)

type ProductService struct {
	Repo interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) CreateProduct(p *models.Product) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return s.Repo.CreateProduct(context.Background(), p)
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.Repo.GetProductByID(context.Background(), id)
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	return s.Repo.GetAllProducts(context.Background())
}

func (s *ProductService) GetAllProductsWithFilters(filters map[string]interface{}) ([]*models.Product, int64, error) {
	return s.Repo.GetAllProductsWithFilters(context.Background(), filters)
}

func (s *ProductService) SearchProducts(query string, page, limit int) ([]*models.Product, int64, error) {
	// Clean and prepare search query
	query = strings.TrimSpace(query)
	if query == "" {
		return []*models.Product{}, 0, nil
	}

	// Create search filters
	searchFilters := map[string]interface{}{
		"search": query,
		"page":   page,
		"limit":  limit,
	}

	return s.Repo.SearchProducts(context.Background(), searchFilters)
}

func (s *ProductService) GetProductsByCategory(category string, page, limit int) ([]*models.Product, int64, error) {
	filters := map[string]interface{}{
		"category": category,
		"page":     page,
		"limit":    limit,
	}

	return s.Repo.GetProductsByCategory(context.Background(), filters)
}

func (s *ProductService) GetFeaturedProducts(limit int) ([]*models.Product, error) {
	return s.Repo.GetFeaturedProducts(context.Background(), limit)
}

func (s *ProductService) GetProductCategories() ([]string, error) {
	return s.Repo.GetProductCategories(context.Background())
}

func (s *ProductService) UpdateProduct(id string, updates map[string]interface{}) error {
	updates["updatedAt"] = time.Now()
	return s.Repo.UpdateProduct(context.Background(), id, updates)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.Repo.DeleteProduct(context.Background(), id)
}

// GetProductsByPriceRange returns products within a specific price range
func (s *ProductService) GetProductsByPriceRange(minPrice, maxPrice float64, page, limit int) ([]*models.Product, int64, error) {
	filters := map[string]interface{}{
		"minPrice": minPrice,
		"maxPrice": maxPrice,
		"page":     page,
		"limit":    limit,
	}

	return s.Repo.GetProductsByPriceRange(context.Background(), filters)
}

// GetProductsInStock returns products that are in stock
func (s *ProductService) GetProductsInStock(page, limit int) ([]*models.Product, int64, error) {
	filters := map[string]interface{}{
		"inStock": true,
		"page":    page,
		"limit":   limit,
	}

	return s.Repo.GetProductsInStock(context.Background(), filters)
}

// GetNewArrivals returns recently added products
func (s *ProductService) GetNewArrivals(limit int) ([]*models.Product, error) {
	return s.Repo.GetNewArrivals(context.Background(), limit)
}

// GetBestSellers returns best selling products
func (s *ProductService) GetBestSellers(limit int) ([]*models.Product, error) {
	return s.Repo.GetBestSellers(context.Background(), limit)
}

// UpdateProductStock updates the stock quantity of a product
func (s *ProductService) UpdateProductStock(productID string, quantity int) error {
	updates := map[string]interface{}{
		"inStock": quantity,
		"updatedAt": time.Now(),
	}
	return s.Repo.UpdateProduct(context.Background(), productID, updates)
}

// GetRelatedProducts returns products related to a given product
func (s *ProductService) GetRelatedProducts(productID string, limit int) ([]*models.Product, error) {
	return s.Repo.GetRelatedProducts(context.Background(), productID, limit)
}
