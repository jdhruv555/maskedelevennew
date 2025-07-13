package middleware

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/Shrey-Yash/Masked11/internal/database"
	"github.com/Shrey-Yash/Masked11/internal/utils"
	"github.com/redis/go-redis/v9"
)

// RateLimiterConfig holds rate limiter configuration
type RateLimiterConfig struct {
	MaxRequests     int           `json:"max_requests"`
	WindowDuration  time.Duration `json:"window_duration"`
	BurstLimit      int           `json:"burst_limit"`
	EnableBurst     bool          `json:"enable_burst"`
	SkipSuccessful  bool          `json:"skip_successful"`
	SkipFailed      bool          `json:"skip_failed"`
	KeyGenerator    func(c *fiber.Ctx) string
	LimitReached    func(c *fiber.Ctx) error
	Storage         limiter.Storage
}

// DefaultRateLimiterConfig returns default rate limiter configuration
func DefaultRateLimiterConfig() *RateLimiterConfig {
	return &RateLimiterConfig{
		MaxRequests:    100,
		WindowDuration: 1 * time.Minute,
		BurstLimit:     10,
		EnableBurst:    true,
		SkipSuccessful: false,
		SkipFailed:     false,
		KeyGenerator:   defaultKeyGenerator,
		LimitReached:   defaultLimitReached,
		Storage:        nil, // Will use memory storage by default
	}
}

// RateLimiter returns a rate limiting middleware
func RateLimiter() fiber.Handler {
	config := DefaultRateLimiterConfig()
	return createRateLimiter(config)
}

// AdaptiveRateLimiter returns a rate limiting middleware with adaptive limits
func AdaptiveRateLimiter() fiber.Handler {
	config := DefaultRateLimiterConfig()
	config.MaxRequests = 200
	config.WindowDuration = 1 * time.Minute
	config.BurstLimit = 20
	return createRateLimiter(config)
}

// StrictRateLimiter returns a strict rate limiting middleware for sensitive endpoints
func StrictRateLimiter() fiber.Handler {
	config := DefaultRateLimiterConfig()
	config.MaxRequests = 10
	config.WindowDuration = 1 * time.Minute
	config.BurstLimit = 2
	return createRateLimiter(config)
}

// AuthRateLimiter returns a rate limiting middleware for authentication endpoints
func AuthRateLimiter() fiber.Handler {
	config := DefaultRateLimiterConfig()
	config.MaxRequests = 5
	config.WindowDuration = 15 * time.Minute
	config.BurstLimit = 1
	return createRateLimiter(config)
}

// createRateLimiter creates a rate limiter with the given configuration
func createRateLimiter(config *RateLimiterConfig) fiber.Handler {
	limiterConfig := limiter.Config{
		Max:        config.MaxRequests,
		Expiration: config.WindowDuration,
		KeyGenerator: func(c *fiber.Ctx) string {
			if config.KeyGenerator != nil {
				return config.KeyGenerator(c)
			}
			return defaultKeyGenerator(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			if config.LimitReached != nil {
				return config.LimitReached(c)
			}
			return defaultLimitReached(c)
		},
		Storage: config.Storage,
		SkipSuccessfulRequests: config.SkipSuccessful,
		SkipFailedRequests:     config.SkipFailed,
	}

	limiter := limiter.New(limiterConfig)

	return func(c *fiber.Ctx) error {
		// Skip rate limiting for health checks and metrics
		if c.Path() == "/health" || c.Path() == "/metrics" {
			return c.Next()
		}

		// Apply rate limiting
		if err := limiter(c); err != nil {
			return err
		}

		// Track rate limiting metrics
		trackRateLimitMetrics(c)

		return c.Next()
	}
}

// defaultKeyGenerator generates a key for rate limiting based on IP and user ID
func defaultKeyGenerator(c *fiber.Ctx) string {
	// Get client IP
	ip := c.IP()
	if ip == "" {
		ip = "unknown"
	}

	// Get user ID if available
	userID := c.Locals("userID")
	if userID != nil {
		return fmt.Sprintf("rate_limit:%s:%s", ip, userID.(string))
	}

	// Get session key if available
	sessionKey := c.Locals("sessionKey")
	if sessionKey != nil {
		return fmt.Sprintf("rate_limit:%s:%s", ip, sessionKey.(string))
	}

	return fmt.Sprintf("rate_limit:%s", ip)
}

// defaultLimitReached handles rate limit exceeded
func defaultLimitReached(c *fiber.Ctx) error {
	// Track rate limit exceeded
	trackRateLimitExceeded(c)

	// Return standardized error response
	return utils.SendErrorResponse(c, fiber.StatusTooManyRequests, "Rate Limit Exceeded", 
		"Too many requests. Please try again later.")
}

// trackRateLimitMetrics tracks rate limiting metrics
func trackRateLimitMetrics(c *fiber.Ctx) {
	// Increment request count
	utils.IncrementRequestCount()

	// Track by endpoint
	endpoint := c.Path()
	utils.IncrementEndpointRequestCount(endpoint)

	// Track by IP
	ip := c.IP()
	utils.IncrementIPRequestCount(ip)
}

// trackRateLimitExceeded tracks when rate limits are exceeded
func trackRateLimitExceeded(c *fiber.Ctx) {
	// Increment rate limit exceeded count
	utils.IncrementRateLimitExceededCount()

	// Track by endpoint
	endpoint := c.Path()
	utils.IncrementEndpointRateLimitExceededCount(endpoint)

	// Track by IP
	ip := c.IP()
	utils.IncrementIPRateLimitExceededCount(ip)

	// Log rate limit exceeded
	utils.LogRateLimitExceeded(c)
}

// RedisRateLimiter returns a rate limiter using Redis storage
func RedisRateLimiter() fiber.Handler {
	config := DefaultRateLimiterConfig()
	config.Storage = &RedisStorage{
		Client: database.Redis,
		Ctx:    database.Ctx,
	}
	return createRateLimiter(config)
}

// RedisStorage implements limiter.Storage interface using Redis
type RedisStorage struct {
	Client *redis.Client
	Ctx    context.Context
}

// Get retrieves the current count for a key
func (rs *RedisStorage) Get(key string) (int, error) {
	val, err := rs.Client.Get(rs.Ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}
	return strconv.Atoi(val)
}

// Set sets the count for a key with expiration
func (rs *RedisStorage) Set(key string, count int, expiration time.Duration) error {
	return rs.Client.Set(rs.Ctx, key, count, expiration).Err()
}

// Delete deletes a key
func (rs *RedisStorage) Delete(key string) error {
	return rs.Client.Del(rs.Ctx, key).Err()
}

// Reset resets the count for a key
func (rs *RedisStorage) Reset(key string) error {
	return rs.Client.Del(rs.Ctx, key).Err()
}

// GetTTL gets the time to live for a key
func (rs *RedisStorage) GetTTL(key string) (time.Duration, error) {
	return rs.Client.TTL(rs.Ctx, key).Result()
}

// Increment increments the count for a key
func (rs *RedisStorage) Increment(key string) error {
	return rs.Client.Incr(rs.Ctx, key).Err()
}

// Decrement decrements the count for a key
func (rs *RedisStorage) Decrement(key string) error {
	return rs.Client.Decr(rs.Ctx, key).Err()
}

// Clear clears all keys
func (rs *RedisStorage) Clear() error {
	// This is a destructive operation, so we'll implement it carefully
	// For now, we'll return an error to prevent accidental clearing
	return fmt.Errorf("clear operation not implemented for safety")
}

// Keys gets all keys with a pattern
func (rs *RedisStorage) Keys(pattern string) ([]string, error) {
	return rs.Client.Keys(rs.Ctx, pattern).Result()
}

// Len gets the number of keys
func (rs *RedisStorage) Len() (int, error) {
	keys, err := rs.Client.Keys(rs.Ctx, "*").Result()
	if err != nil {
		return 0, err
	}
	return len(keys), nil
} 