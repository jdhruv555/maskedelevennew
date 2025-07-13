package postgres

import (
	"context"
	"errors"
	"time"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/repositories/interfaces"
	"github.com/Shrey-Yash/Masked11/internal/constants"
)

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) interfaces.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order *models.Order, items []models.OrderItem) error {
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO orders (id, user_id, total, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = tx.Exec(ctx, query, order.ID, order.UserID, order.Total, order.Status, time.Now(), time.Now())
	if err != nil {
		return err
	}

	itemQuery := `INSERT INTO order_items (id, order_id, product_id, name, price, quantity, image, size, subtotal) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	for _, item := range items {
		_, err := tx.Exec(ctx, itemQuery, item.ID, order.ID, item.ProductID, item.Name, item.Price, item.Quantity, item.Image, item.Size, item.Subtotal)
		if err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func (r *orderRepository) GetOrdersByUserID(ctx context.Context, userID string) ([]models.Order, error) {
	orders := []models.Order{}

	rows, err := r.db.Query(ctx, `SELECT id, user_id, total, status, created_at, updated_at FROM orders WHERE user_id = $1 ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.Total, &order.Status, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		
		items, err := r.getOrderItems(ctx, order.ID.String())
		if err != nil {
			return nil, err
		}
		order.Items = items
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) GetOrderByID(ctx context.Context, orderID string) (*models.Order, error) {
	var order models.Order
	err := r.db.QueryRow(ctx, `SELECT id, user_id, total, status, created_at, updated_at FROM orders WHERE id = $1`, orderID).Scan(
		&order.ID, &order.UserID, &order.Total, &order.Status, &order.CreatedAt, &order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	items, err := r.getOrderItems(ctx, orderID)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return &order, nil
}

func (r *orderRepository) UpdateOrderStatus(orderID string, status string) error {

	normalized := strings.ToLower(status)

	var validStatus string
	switch normalized {
	case strings.ToLower(constants.OrderStatusPending):
		validStatus = constants.OrderStatusPending
	case strings.ToLower(constants.OrderStatusShipped):
		validStatus = constants.OrderStatusShipped
	case strings.ToLower(constants.OrderStatusDelivered):
		validStatus = constants.OrderStatusDelivered
	case strings.ToLower(constants.OrderStatusCancelled):
		validStatus = constants.OrderStatusCancelled
	default:
		return errors.New("invalid order status")
	}

	cmd, err := r.db.Exec(context.Background(), `UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3`, validStatus, time.Now(), orderID)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

func (r *orderRepository) DeleteOrder(orderID string) error {
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `DELETE FROM order_items WHERE order_id = $1`, orderID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, `DELETE FROM orders WHERE id = $1`, orderID)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *orderRepository) CancelOrder(orderID string) error {
	cmd, err := r.db.Exec(context.Background(), `UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3`, constants.OrderStatusCancelled, time.Now(), orderID)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

func (r *orderRepository) getOrderItems(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	items := []models.OrderItem{}

	rows, err := r.db.Query(ctx, `SELECT id, order_id, product_id, name, price, quantity, image, size, subtotal FROM order_items WHERE order_id = $1`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Name, &item.Price, &item.Quantity, &item.Image, &item.Size, &item.Subtotal); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
