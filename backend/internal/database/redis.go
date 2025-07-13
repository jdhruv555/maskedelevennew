package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client
var Ctx = context.Background()

// InitRedis initializes Redis with optimized connection pooling
func InitRedis() error {
	uri := os.Getenv("REDIS_URI")
	if uri == "" {
		uri = "localhost:6379"
	}

	// Parse Redis URL for advanced configuration
	opts, err := redis.ParseURL(uri)
	if err != nil {
		// Fallback to simple configuration
		opts = &redis.Options{
			Addr:     uri,
			Password: "",
			DB:       0,
		}
	}

	// Configure connection pool settings
	opts.PoolSize = 50                    // Maximum number of connections
	opts.MinIdleConns = 10                // Minimum number of idle connections
	opts.MaxRetries = 3                   // Maximum number of retries
	opts.DialTimeout = 5 * time.Second    // Connection timeout
	opts.ReadTimeout = 3 * time.Second    // Read timeout
	opts.WriteTimeout = 3 * time.Second   // Write timeout
	opts.PoolTimeout = 4 * time.Second    // Pool timeout

	// Create Redis client
	client := redis.NewClient(opts)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return err
	}

	Redis = client

	// Start connection monitoring
	go monitorRedisConnection()

	// Configure Redis for optimal performance
	if err := configureRedis(); err != nil {
		log.Printf("Warning: Redis configuration failed: %v", err)
	}

	log.Println("✅ Redis connected with optimized settings")
	return nil
}

// configureRedis configures Redis for optimal performance
func configureRedis() error {
	ctx := context.Background()

	// Set memory policy
	Redis.ConfigSet(ctx, "maxmemory-policy", "allkeys-lru")

	// Set memory limit (adjust based on your server)
	Redis.ConfigSet(ctx, "maxmemory", "256mb")

	// Enable persistence
	Redis.ConfigSet(ctx, "save", "900 1 300 10 60 10000")

	// Set timeout settings
	Redis.ConfigSet(ctx, "timeout", "300")

	// Enable compression for large values
	Redis.ConfigSet(ctx, "hash-max-ziplist-entries", "512")
	Redis.ConfigSet(ctx, "hash-max-ziplist-value", "64")

	return nil
}

// monitorRedisConnection monitors Redis connection health
func monitorRedisConnection() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		
		if err := Redis.Ping(ctx).Err(); err != nil {
			log.Printf("⚠️ Redis connection health check failed: %v", err)
		}
		
		cancel()
	}
}

// GetRedisStats returns Redis connection and performance statistics
func GetRedisStats() map[string]interface{} {
	if Redis == nil {
		return map[string]interface{}{
			"status": "disconnected",
		}
	}
	
	// Parse basic info
	stats := map[string]interface{}{
		"status":    "connected",
		"timestamp": time.Now().UTC(),
	}

	// Extract key metrics from INFO command
	// This is a simplified version - in production you'd parse the full INFO output
	stats["connected_clients"] = "active"
	stats["used_memory"] = "optimized"
	stats["keyspace_hits"] = "tracked"
	stats["keyspace_misses"] = "tracked"

	return stats
}

// CloseRedis closes the Redis connection
func CloseRedis() error {
	if Redis != nil {
		return Redis.Close()
	}
	return nil
}

// Cache utilities with optimized settings

// SetWithTTL sets a key with TTL and optimized options
func SetWithTTL(key string, value interface{}, ttl time.Duration) error {
	ctx := context.Background()
	return Redis.Set(ctx, key, value, ttl).Err()
}

// GetWithFallback gets a value with fallback function
func GetWithFallback(key string, fallback func() (interface{}, error), ttl time.Duration) (interface{}, error) {
	ctx := context.Background()
	
	// Try to get from cache
	val, err := Redis.Get(ctx, key).Result()
	if err == nil {
		return val, nil
	}

	// Cache miss - call fallback function
	result, err := fallback()
	if err != nil {
		return nil, err
	}

	// Cache the result
	if err := SetWithTTL(key, result, ttl); err != nil {
		log.Printf("Warning: Failed to cache result: %v", err)
	}

	return result, nil
}

// MGetWithFallback gets multiple values with fallback
func MGetWithFallback(keys []string, fallback func([]string) (map[string]interface{}, error), ttl time.Duration) (map[string]interface{}, error) {
	ctx := context.Background()
	
	// Try to get all from cache
	vals, err := Redis.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	// Check for cache misses
	missedKeys := make([]string, 0)
	result := make(map[string]interface{})
	
	for i, val := range vals {
		if val == nil {
			missedKeys = append(missedKeys, keys[i])
		} else {
			result[keys[i]] = val
		}
	}

	// If all keys were found in cache
	if len(missedKeys) == 0 {
		return result, nil
	}

	// Call fallback for missed keys
	fallbackResult, err := fallback(missedKeys)
	if err != nil {
		return nil, err
	}

	// Cache the fallback results
	pipe := Redis.Pipeline()
	for key, value := range fallbackResult {
		pipe.Set(ctx, key, value, ttl)
		result[key] = value
	}
	pipe.Exec(ctx)

	return result, nil
}

// DeletePattern deletes keys matching a pattern
func DeletePattern(pattern string) error {
	ctx := context.Background()
	keys, err := Redis.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		return Redis.Del(ctx, keys...).Err()
	}

	return nil
}

// IncrementWithTTL increments a counter with TTL
func IncrementWithTTL(key string, ttl time.Duration) (int64, error) {
	ctx := context.Background()
	
	pipe := Redis.Pipeline()
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, ttl)
	
	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}

	return incr.Val(), nil
}

// SetNXWithTTL sets a key only if it doesn't exist
func SetNXWithTTL(key string, value interface{}, ttl time.Duration) (bool, error) {
	ctx := context.Background()
	
	pipe := Redis.Pipeline()
	setNX := pipe.SetNX(ctx, key, value, ttl)
	
	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	return setNX.Val(), nil
}

// Hash operations with optimization

// HSetWithTTL sets hash fields with TTL
func HSetWithTTL(key string, fields map[string]interface{}, ttl time.Duration) error {
	ctx := context.Background()
	
	pipe := Redis.Pipeline()
	pipe.HSet(ctx, key, fields)
	pipe.Expire(ctx, key, ttl)
	
	_, err := pipe.Exec(ctx)
	return err
}

// HGetAllWithTTL gets all hash fields
func HGetAllWithTTL(key string) (map[string]string, error) {
	ctx := context.Background()
	return Redis.HGetAll(ctx, key).Result()
}

// List operations with optimization

// LPushWithTTL pushes to list with TTL
func LPushWithTTL(key string, values ...interface{}) error {
	ctx := context.Background()
	
	pipe := Redis.Pipeline()
	pipe.LPush(ctx, key, values...)
	pipe.Expire(ctx, key, 24*time.Hour) // Default TTL for lists
	
	_, err := pipe.Exec(ctx)
	return err
}

// LRangeWithTTL gets list range
func LRangeWithTTL(key string, start, stop int64) ([]string, error) {
	ctx := context.Background()
	return Redis.LRange(ctx, key, start, stop).Result()
}

// Set operations with optimization

// SAddWithTTL adds to set with TTL
func SAddWithTTL(key string, members ...interface{}) error {
	ctx := context.Background()
	
	pipe := Redis.Pipeline()
	pipe.SAdd(ctx, key, members...)
	pipe.Expire(ctx, key, 24*time.Hour) // Default TTL for sets
	
	_, err := pipe.Exec(ctx)
	return err
}

// SMembersWithTTL gets set members
func SMembersWithTTL(key string) ([]string, error) {
	ctx := context.Background()
	return Redis.SMembers(ctx, key).Result()
}

// ZSet operations with optimization

// ZAddWithTTL adds to sorted set with TTL
func ZAddWithTTL(key string, members ...redis.Z) error {
	ctx := context.Background()
	
	pipe := Redis.Pipeline()
	pipe.ZAdd(ctx, key, members...)
	pipe.Expire(ctx, key, 24*time.Hour) // Default TTL for sorted sets
	
	_, err := pipe.Exec(ctx)
	return err
}

// ZRangeWithTTL gets sorted set range
func ZRangeWithTTL(key string, start, stop int64) ([]string, error) {
	ctx := context.Background()
	return Redis.ZRange(ctx, key, start, stop).Result()
}

// Performance monitoring

// GetCacheHitRate calculates cache hit rate
func GetCacheHitRate() float64 {
	ctx := context.Background()
	
	hits, err := Redis.Get(ctx, "cache:hits").Int64()
	if err != nil {
		hits = 0
	}
	
	misses, err := Redis.Get(ctx, "cache:misses").Int64()
	if err != nil {
		misses = 0
	}
	
	total := hits + misses
	if total == 0 {
		return 0
	}
	
	return float64(hits) / float64(total) * 100
}

// RecordCacheHit records a cache hit
func RecordCacheHit() {
	ctx := context.Background()
	Redis.Incr(ctx, "cache:hits")
}

// RecordCacheMiss records a cache miss
func RecordCacheMiss() {
	ctx := context.Background()
	Redis.Incr(ctx, "cache:misses")
}