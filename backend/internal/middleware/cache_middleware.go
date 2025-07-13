package middleware

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/Shrey-Yash/Masked11/internal/database"
)

// CacheMiddleware provides intelligent caching for API responses
func CacheMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip caching for non-GET requests
		if c.Method() != "GET" {
			return c.Next()
		}

		// Skip caching for authenticated requests
		if c.Get("Authorization") != "" {
			return c.Next()
		}

		// Generate cache key based on request
		cacheKey := generateCacheKey(c)

		// Try to get from cache
		cached, err := database.Redis.Get(database.Ctx, cacheKey).Result()
		if err == nil {
			// Cache hit - return cached response
			var response map[string]interface{}
			if err := json.Unmarshal([]byte(cached), &response); err == nil {
				c.Set("X-Cache", "HIT")
				c.Set("X-Cache-Key", cacheKey)
				return c.JSON(response)
			}
		}

		// Cache miss - store original response
		originalBody := c.Response().Body()
		originalStatus := c.Response().StatusCode()

		// Continue to handler
		if err := c.Next(); err != nil {
			return err
		}

		// Cache successful responses only
		if c.Response().StatusCode() == fiber.StatusOK {
			response := map[string]interface{}{
				"data":    json.RawMessage(c.Response().Body()),
				"status":  c.Response().StatusCode(),
				"cached":  true,
				"cachedAt": time.Now().UTC(),
			}

			// Serialize response
			responseBytes, err := json.Marshal(response)
			if err == nil {
				// Set cache with TTL based on endpoint
				ttl := getCacheTTL(c.Path())
				database.Redis.Set(database.Ctx, cacheKey, responseBytes, ttl)
			}
		}

		c.Set("X-Cache", "MISS")
		c.Set("X-Cache-Key", cacheKey)
		return nil
	}
}

// generateCacheKey creates a unique cache key based on request parameters
func generateCacheKey(c *fiber.Ctx) string {
	// Include path and query parameters
	key := fmt.Sprintf("cache:%s:%s", c.Method(), c.Path())
	
	// Add query parameters to cache key
	query := c.Query("page") + c.Query("limit") + c.Query("category") + 
		c.Query("search") + c.Query("sortBy") + c.Query("minPrice") + 
		c.Query("maxPrice")

	if query != "" {
		hash := md5.Sum([]byte(query))
		key += fmt.Sprintf(":%x", hash)
	}

	return key
}

// getCacheTTL returns appropriate TTL based on endpoint
func getCacheTTL(path string) time.Duration {
	switch {
	case path == "/api/products/categories":
		return 24 * time.Hour // Categories change rarely
	case path == "/api/products/featured":
		return 1 * time.Hour // Featured products update periodically
	case path == "/api/products":
		return 30 * time.Minute // Product list with filters
	default:
		return 15 * time.Minute // Default cache time
	}
}

// CacheInvalidationMiddleware invalidates cache when data changes
func CacheInvalidationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Continue to handler first
		if err := c.Next(); err != nil {
			return err
		}

		// Invalidate cache for write operations
		if c.Method() == "POST" || c.Method() == "PUT" || c.Method() == "DELETE" {
			go invalidateRelatedCache(c.Path())
		}

		return nil
	}
}

// invalidateRelatedCache removes related cache entries
func invalidateRelatedCache(path string) {
	ctx := context.Background()
	
	// Define cache patterns to invalidate
	patterns := []string{
		"cache:GET:/api/products*",
		"cache:GET:/api/products/categories*",
		"cache:GET:/api/products/featured*",
	}

	for _, pattern := range patterns {
		keys, err := database.Redis.Keys(ctx, pattern).Result()
		if err == nil {
			for _, key := range keys {
				database.Redis.Del(ctx, key)
			}
		}
	}
}

// CacheStatsMiddleware provides cache statistics
func CacheStatsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		
		if err := c.Next(); err != nil {
			return err
		}

		// Add cache statistics to response headers
		cacheHit := c.Get("X-Cache") == "HIT"
		if cacheHit {
			c.Set("X-Cache-Hit", "true")
			c.Set("X-Response-Time", time.Since(start).String())
		}

		return nil
	}
} 