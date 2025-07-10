package handlers

import (
	"encoding/json"
	"log"
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
	products, err := h.Service.GetAllProducts()
	if err != nil {
		log.Println("GetAllProducts error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch products")
	}

	return c.JSON(products)
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
