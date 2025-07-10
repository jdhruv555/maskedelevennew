package redisrepo

import (
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type SessionRepository struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewSessionRepository(client *redis.Client, ctx context.Context) *SessionRepository {
	return &SessionRepository{
		Client: client,
		Ctx:    ctx,
	}
}

func (r *SessionRepository) SaveSession(jti, userID string, duration time.Duration) error {
	return r.Client.Set(r.Ctx, jti, userID, duration).Err()
}

func (r *SessionRepository) DeleteSession(jti string) error {
	return r.Client.Del(r.Ctx, jti).Err()
}

func (r *SessionRepository) StoreSession(jti string, userID string, ttl time.Duration) error {
	return r.Client.Set(r.Ctx, jti, userID, ttl).Err()
}

func (r *SessionRepository) IsSessionActive(jti string) (bool, error) {
	val, err := r.Client.Get(r.Ctx, jti).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return val != "", nil
}
