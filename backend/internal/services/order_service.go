package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/Shrey-Yash/Masked11/internal/constants"
	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/repositories/interfaces"
)

type OrderService struct {
	OrderRepo interfaces.OrderRepository
	CartRepo  interfaces.CartRepository
	UserRepo interfaces.UserRepository
}

func NewOrderService(orderRepo interfaces.OrderRepository, cartRepo interfaces.CartRepository, userRepo interfaces.UserRepository) *OrderService {
	return &OrderService{
		OrderRepo: orderRepo, 
		CartRepo: cartRepo,
		UserRepo: userRepo,
	}
}

func (s *OrderService) CreateOrder(userID string) error {
	cart, err := s.CartRepo.GetCart(userID)
	if err != nil || cart == nil || len(cart.Items) == 0 {
		if err != nil {
			fmt.Println("CartRepo error:", err)
		} else {
			fmt.Println("Empty or nil cart")
		}
		return errors.New("invalid cart")
	}

	orderID := uuid.New()
	var orderItems []models.OrderItem
	total := 0.0

	for _, item := range cart.Items {
		sub := float64(item.Quantity) * item.Price
		orderItems = append(orderItems, models.OrderItem{
			ID:        uuid.New(),
			OrderID:   orderID,
			ProductID: item.ProductID,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  item.Quantity,
			Size:      item.Size,
			Image:     item.Image,
			Subtotal:  sub,
		})
		total += sub
	}

	order := &models.Order{
		ID:        orderID,
		UserID:    userID,
		Status:    constants.OrderStatusPending,
		Total:     total,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.OrderRepo.CreateOrder(order, orderItems); err != nil {
		return err
	}

	_ = s.CartRepo.DeleteCart(userID)
	return nil
}

func (s *OrderService) GetOrdersByUserID(ctx context.Context, userID string) ([]models.Order, error) {
	return s.OrderRepo.GetOrdersByUserID(ctx, userID)
}

func (s *OrderService) GetOrderByID(ctx context.Context, orderID string, isAdmin bool, userID string) (*models.Order, error) {
	order, err := s.OrderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) UpdateOrderStatus(orderID, status string) error {
	return s.OrderRepo.UpdateOrderStatus(orderID, status)
}

func (s *OrderService) DeleteOrder(orderID string) error {
	return s.OrderRepo.DeleteOrder(orderID)
}

func (s *OrderService) CancelOrder(ctx context.Context, userID string, orderID string) error {
	order, err := s.OrderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil
	}

	if order.UserID != userID {
		return errors.New("unauthorized access to cancel order")
	}
	return s.OrderRepo.CancelOrder(orderID)
}
