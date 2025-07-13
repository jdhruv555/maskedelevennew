package handlers

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/services"
)

type ProductHandler struct {
	Service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := json.Unmarshal(c.Body(), &product); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON body")
	}

	err := h.Service.CreateProduct(&product)
	if err != nil {
		log.Println("CreateProduct error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create product")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": product,
	})
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Product ID required")
	}

	product, err := h.Service.GetProductByID(id)
	if err != nil {
		log.Println("GetProductByID error:", err)
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	return c.JSON(product)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	// Get query parameters for filtering and pagination
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "12"))
	search := c.Query("search", "")
	category := c.Query("category", "")
	minPrice, _ := strconv.ParseFloat(c.Query("minPrice", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(c.Query("maxPrice", "0"), 64)
	sortBy := c.Query("sortBy", "createdAt")
	sortOrder := c.Query("sortOrder", "desc")

	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 12
	}

	// Create filter options
	filters := map[string]interface{}{
		"search":     search,
		"category":   category,
		"minPrice":   minPrice,
		"maxPrice":   maxPrice,
		"sortBy":     sortBy,
		"sortOrder":  sortOrder,
		"page":       page,
		"limit":      limit,
	}

	products, total, err := h.Service.GetAllProductsWithFilters(filters)
	if err != nil {
		log.Println("GetAllProducts error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch products")
	}

	// Calculate pagination info
	totalPages := (total + limit - 1) / limit
	hasNext := page < totalPages
	hasPrev := page > 1

	return c.JSON(fiber.Map{
		"products":   products,
		"pagination": fiber.Map{
			"currentPage": page,
			"totalPages":  totalPages,
			"totalItems":  total,
			"limit":       limit,
			"hasNext":     hasNext,
			"hasPrev":     hasPrev,
		},
		"filters": filters,
	})
}

func (h *ProductHandler) SearchProducts(c *fiber.Ctx) error {
	query := c.Query("q", "")
	if query == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Search query required")
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "12"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 12
	}

	products, total, err := h.Service.SearchProducts(query, page, limit)
	if err != nil {
		log.Println("SearchProducts error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to search products")
	}

	totalPages := (total + limit - 1) / limit

	return c.JSON(fiber.Map{
		"products":   products,
		"query":      query,
		"pagination": fiber.Map{
			"currentPage": page,
			"totalPages":  totalPages,
			"totalItems":  total,
			"limit":       limit,
		},
	})
}

func (h *ProductHandler) GetProductsByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	if category == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Category required")
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "12"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 12
	}

	products, total, err := h.Service.GetProductsByCategory(category, page, limit)
	if err != nil {
		log.Println("GetProductsByCategory error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch products by category")
	}

	totalPages := (total + limit - 1) / limit

	return c.JSON(fiber.Map{
		"products":   products,
		"category":   category,
		"pagination": fiber.Map{
			"currentPage": page,
			"totalPages":  totalPages,
			"totalItems":  total,
			"limit":       limit,
		},
	})
}

func (h *ProductHandler) GetFeaturedProducts(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "8"))
	if limit < 1 || limit > 20 {
		limit = 8
	}

	products, err := h.Service.GetFeaturedProducts(limit)
	if err != nil {
		log.Println("GetFeaturedProducts error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch featured products")
	}

	return c.JSON(fiber.Map{
		"products": products,
	})
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Product ID required")
	}

	var updates map[string]interface{}
	if err := json.Unmarshal(c.Body(), &updates); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON body")
	}

	delete(updates, "_id")
	updates["updatedAt"] = time.Now()

	err := h.Service.UpdateProduct(id, updates)
	if err != nil {
		log.Println("UpdateProduct error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update product")
	}

	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Product ID required")
	}

	err := h.Service.DeleteProduct(id)
	if err != nil {
		log.Println("DeleteProduct error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete product")
	}

	return c.JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func (h *ProductHandler) GetProductCategories(c *fiber.Ctx) error {
	categories, err := h.Service.GetProductCategories()
	if err != nil {
		log.Println("GetProductCategories error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch categories")
	}

	return c.JSON(fiber.Map{
		"categories": categories,
	})
}
