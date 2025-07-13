package redisrepo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	
	"github.com/Shrey-Yash/Masked11/internal/models"
	"github.com/Shrey-Yash/Masked11/internal/repositories/interfaces"
)

type cartRepository struct {
	rdb *redis.Client
	ctx context.Context
}

func NewCartRepository(rdb *redis.Client, ctx context.Context) interfaces.CartRepository {
	return &cartRepository{rdb: rdb, ctx: ctx}
}

func (r *cartRepository) GetCart(key string) (*models.Cart, error) {
	data, err := r.rdb.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var cart models.Cart
	if err := json.Unmarshal([]byte(data), &cart); err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepository) SetCart(key string, cart *models.Cart) error {
	cart.UpdatedAt = time.Now()
	data, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	return r.rdb.Set(r.ctx, key, data, 0).Err()
}

func (r *cartRepository) DeleteCart(key string) error {
	return r.rdb.Del(r.ctx, key).Err()
}