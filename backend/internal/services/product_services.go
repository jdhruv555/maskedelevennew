package services

import (
	"context"
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

func (s *ProductService) UpdateProduct(id string, updates map[string]interface{}) error {
	updates["updatedAt"] = time.Now()
	return s.Repo.UpdateProduct(context.Background(), id, updates)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.Repo.DeleteProduct(context.Background(), id)
}
