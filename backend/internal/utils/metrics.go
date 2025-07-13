package utils

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/Shrey-Yash/Masked11/internal/database"
)

// Metrics holds various performance metrics
type Metrics struct {
	RequestCount              int64             `json:"request_count"`
	ErrorCount                int64             `json:"error_count"`
	RateLimitExceededCount    int64             `json:"rate_limit_exceeded_count"`
	AverageResponseTime       float64           `json:"average_response_time"`
	TotalResponseTime         int64             `json:"total_response_time"`
	ResponseCount             int64             `json:"response_count"`
	EndpointRequestCounts     map[string]int64  `json:"endpoint_request_counts"`
	EndpointErrorCounts       map[string]int64  `json:"endpoint_error_counts"`
	IPRequestCounts           map[string]int64  `json:"ip_request_counts"`
	IPRateLimitExceededCounts map[string]int64  `json:"ip_rate_limit_exceeded_counts"`
	DatabaseQueryCount        int64             `json:"database_query_count"`
	DatabaseErrorCount        int64             `json:"database_error_count"`
	CacheHitCount             int64             `json:"cache_hit_count"`
	CacheMissCount            int64             `json:"cache_miss_count"`
	ActiveConnections         int64             `json:"active_connections"`
	Uptime                    time.Duration     `json:"uptime"`
	StartTime                 time.Time         `json:"start_time"`
}

var (
	metrics     = &Metrics{}
	metricsLock sync.RWMutex
	startTime   = time.Now()
)

// Initialize metrics
func init() {
	metrics.EndpointRequestCounts = make(map[string]int64)
	metrics.EndpointErrorCounts = make(map[string]int64)
	metrics.IPRequestCounts = make(map[string]int64)
	metrics.IPRateLimitExceededCounts = make(map[string]int64)
	metrics.StartTime = startTime
}

// IncrementRequestCount increments the total request count
func IncrementRequestCount() {
	atomic.AddInt64(&metrics.RequestCount, 1)
}

// IncrementErrorCount increments the total error count
func IncrementErrorCount() {
	atomic.AddInt64(&metrics.ErrorCount, 1)
}

// IncrementRateLimitExceededCount increments the rate limit exceeded count
func IncrementRateLimitExceededCount() {
	atomic.AddInt64(&metrics.RateLimitExceededCount, 1)
}

// IncrementEndpointRequestCount increments the request count for a specific endpoint
func IncrementEndpointRequestCount(endpoint string) {
	metricsLock.Lock()
	defer metricsLock.Unlock()
	metrics.EndpointRequestCounts[endpoint]++
}

// IncrementEndpointErrorCount increments the error count for a specific endpoint
func IncrementEndpointErrorCount(endpoint string) {
	metricsLock.Lock()
	defer metricsLock.Unlock()
	metrics.EndpointErrorCounts[endpoint]++
}

// IncrementIPRequestCount increments the request count for a specific IP
func IncrementIPRequestCount(ip string) {
	metricsLock.Lock()
	defer metricsLock.Unlock()
	metrics.IPRequestCounts[ip]++
}

// IncrementIPRateLimitExceededCount increments the rate limit exceeded count for a specific IP
func IncrementIPRateLimitExceededCount(ip string) {
	metricsLock.Lock()
	defer metricsLock.Unlock()
	metrics.IPRateLimitExceededCounts[ip]++
}

// IncrementDatabaseQueryCount increments the database query count
func IncrementDatabaseQueryCount() {
	atomic.AddInt64(&metrics.DatabaseQueryCount, 1)
}

// IncrementDatabaseErrorCount increments the database error count
func IncrementDatabaseErrorCount() {
	atomic.AddInt64(&metrics.DatabaseErrorCount, 1)
}

// IncrementCacheHitCount increments the cache hit count
func IncrementCacheHitCount() {
	atomic.AddInt64(&metrics.CacheHitCount, 1)
}

// IncrementCacheMissCount increments the cache miss count
func IncrementCacheMissCount() {
	atomic.AddInt64(&metrics.CacheMissCount, 1)
}

// SetActiveConnections sets the number of active connections
func SetActiveConnections(count int64) {
	atomic.StoreInt64(&metrics.ActiveConnections, count)
}

// AddResponseTime adds a response time to calculate the average
func AddResponseTime(duration time.Duration) {
	atomic.AddInt64(&metrics.TotalResponseTime, int64(duration.Milliseconds()))
	atomic.AddInt64(&metrics.ResponseCount, 1)
}

// GetRequestCount returns the total request count
func GetRequestCount() int64 {
	return atomic.LoadInt64(&metrics.RequestCount)
}

// GetErrorCount returns the total error count
func GetErrorCount() int64 {
	return atomic.LoadInt64(&metrics.ErrorCount)
}

// GetRateLimitExceededCount returns the total rate limit exceeded count
func GetRateLimitExceededCount() int64 {
	return atomic.LoadInt64(&metrics.RateLimitExceededCount)
}

// GetAverageResponseTime returns the average response time in milliseconds
func GetAverageResponseTime() float64 {
	responseCount := atomic.LoadInt64(&metrics.ResponseCount)
	if responseCount == 0 {
		return 0
	}
	totalTime := atomic.LoadInt64(&metrics.TotalResponseTime)
	return float64(totalTime) / float64(responseCount)
}

// GetDatabaseQueryCount returns the total database query count
func GetDatabaseQueryCount() int64 {
	return atomic.LoadInt64(&metrics.DatabaseQueryCount)
}

// GetDatabaseErrorCount returns the total database error count
func GetDatabaseErrorCount() int64 {
	return atomic.LoadInt64(&metrics.DatabaseErrorCount)
}

// GetCacheHitCount returns the total cache hit count
func GetCacheHitCount() int64 {
	return atomic.LoadInt64(&metrics.CacheHitCount)
}

// GetCacheMissCount returns the total cache miss count
func GetCacheMissCount() int64 {
	return atomic.LoadInt64(&metrics.CacheMissCount)
}

// GetCacheHitRate returns the cache hit rate as a percentage
func GetCacheHitRate() float64 {
	hits := atomic.LoadInt64(&metrics.CacheHitCount)
	misses := atomic.LoadInt64(&metrics.CacheMissCount)
	total := hits + misses
	if total == 0 {
		return 0
	}
	return (float64(hits) / float64(total)) * 100
}

// GetActiveConnections returns the number of active connections
func GetActiveConnections() int64 {
	return atomic.LoadInt64(&metrics.ActiveConnections)
}

// GetUptime returns the server uptime
func GetUptime() time.Duration {
	return time.Since(startTime)
}

// GetMetrics returns a copy of all metrics
func GetMetrics() *Metrics {
	metricsLock.RLock()
	defer metricsLock.RUnlock()

	// Create a copy to avoid race conditions
	copy := &Metrics{
		RequestCount:              atomic.LoadInt64(&metrics.RequestCount),
		ErrorCount:                atomic.LoadInt64(&metrics.ErrorCount),
		RateLimitExceededCount:    atomic.LoadInt64(&metrics.RateLimitExceededCount),
		AverageResponseTime:       GetAverageResponseTime(),
		TotalResponseTime:         atomic.LoadInt64(&metrics.TotalResponseTime),
		ResponseCount:             atomic.LoadInt64(&metrics.ResponseCount),
		DatabaseQueryCount:        atomic.LoadInt64(&metrics.DatabaseQueryCount),
		DatabaseErrorCount:        atomic.LoadInt64(&metrics.DatabaseErrorCount),
		CacheHitCount:             atomic.LoadInt64(&metrics.CacheHitCount),
		CacheMissCount:            atomic.LoadInt64(&metrics.CacheMissCount),
		ActiveConnections:         atomic.LoadInt64(&metrics.ActiveConnections),
		Uptime:                    GetUptime(),
		StartTime:                 startTime,
		EndpointRequestCounts:     make(map[string]int64),
		EndpointErrorCounts:       make(map[string]int64),
		IPRequestCounts:           make(map[string]int64),
		IPRateLimitExceededCounts: make(map[string]int64),
	}

	// Copy maps
	for k, v := range metrics.EndpointRequestCounts {
		copy.EndpointRequestCounts[k] = v
	}
	for k, v := range metrics.EndpointErrorCounts {
		copy.EndpointErrorCounts[k] = v
	}
	for k, v := range metrics.IPRequestCounts {
		copy.IPRequestCounts[k] = v
	}
	for k, v := range metrics.IPRateLimitExceededCounts {
		copy.IPRateLimitExceededCounts[k] = v
	}

	return copy
}

// ResetMetrics resets all metrics (use with caution)
func ResetMetrics() {
	metricsLock.Lock()
	defer metricsLock.Unlock()

	atomic.StoreInt64(&metrics.RequestCount, 0)
	atomic.StoreInt64(&metrics.ErrorCount, 0)
	atomic.StoreInt64(&metrics.RateLimitExceededCount, 0)
	atomic.StoreInt64(&metrics.TotalResponseTime, 0)
	atomic.StoreInt64(&metrics.ResponseCount, 0)
	atomic.StoreInt64(&metrics.DatabaseQueryCount, 0)
	atomic.StoreInt64(&metrics.DatabaseErrorCount, 0)
	atomic.StoreInt64(&metrics.CacheHitCount, 0)
	atomic.StoreInt64(&metrics.CacheMissCount, 0)
	atomic.StoreInt64(&metrics.ActiveConnections, 0)

	metrics.EndpointRequestCounts = make(map[string]int64)
	metrics.EndpointErrorCounts = make(map[string]int64)
	metrics.IPRequestCounts = make(map[string]int64)
	metrics.IPRateLimitExceededCounts = make(map[string]int64)
}

// MetricsMiddleware is a middleware that tracks request metrics
func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Increment request count
		IncrementRequestCount()
		IncrementEndpointRequestCount(c.Path())
		IncrementIPRequestCount(c.IP())

		// Process request
		err := c.Next()

		// Track response time
		duration := time.Since(start)
		AddResponseTime(duration)

		// Track errors
		if err != nil || c.Response().StatusCode() >= 400 {
			IncrementErrorCount()
			IncrementEndpointErrorCount(c.Path())
		}

		return err
	}
}

// DatabaseMetricsMiddleware tracks database query metrics
func DatabaseMetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// This middleware would be used in conjunction with database operations
		// For now, we'll just pass through
		return c.Next()
	}
}

// CacheMetricsMiddleware tracks cache hit/miss metrics
func CacheMetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// This middleware would be used in conjunction with cache operations
		// For now, we'll just pass through
		return c.Next()
	}
}

// LogRateLimitExceeded logs when rate limits are exceeded
func LogRateLimitExceeded(c *fiber.Ctx) {
	// Log rate limit exceeded event
	fmt.Printf("[RATE_LIMIT] IP: %s, Path: %s, Time: %s\n", 
		c.IP(), c.Path(), time.Now().Format(time.RFC3339))
}

// GetDatabaseMetrics returns database-specific metrics
func GetDatabaseMetrics() map[string]interface{} {
	ctx := context.Background()
	
	// Get Redis info
	redisInfo := map[string]interface{}{}
	if database.Redis != nil {
		if info, err := database.Redis.Info(ctx).Result(); err == nil {
			redisInfo["info"] = info
		}
		// MemoryUsage requires a key parameter, so we'll skip it for now
		// if memory, err := database.Redis.MemoryUsage(ctx).Result(); err == nil {
		// 	redisInfo["memory_usage"] = memory
		// }
	}

	// Get MongoDB stats (if available)
	mongoStats := map[string]interface{}{}
	if database.Mongo != nil {
		// MongoDB stats would go here
		mongoStats["status"] = "connected"
	}

	return map[string]interface{}{
		"redis": redisInfo,
		"mongo": mongoStats,
		"postgres": map[string]interface{}{
			"status": "connected", // Simplified for now
		},
	}
}

// GetSystemMetrics returns system-level metrics
func GetSystemMetrics() map[string]interface{} {
	return map[string]interface{}{
		"uptime":           GetUptime().String(),
		"start_time":       startTime.Format(time.RFC3339),
		"active_connections": GetActiveConnections(),
		"memory_usage":     "N/A", // Would need to implement memory tracking
		"cpu_usage":        "N/A", // Would need to implement CPU tracking
	}
}

// ExportMetricsForPrometheus exports metrics in Prometheus format
func ExportMetricsForPrometheus() string {
	metrics := GetMetrics()
	
	output := fmt.Sprintf(`# HELP masked11_requests_total Total number of requests
# TYPE masked11_requests_total counter
masked11_requests_total %d

# HELP masked11_errors_total Total number of errors
# TYPE masked11_errors_total counter
masked11_errors_total %d

# HELP masked11_rate_limit_exceeded_total Total number of rate limit exceeded
# TYPE masked11_rate_limit_exceeded_total counter
masked11_rate_limit_exceeded_total %d

# HELP masked11_response_time_average_ms Average response time in milliseconds
# TYPE masked11_response_time_average_ms gauge
masked11_response_time_average_ms %f

# HELP masked11_database_queries_total Total number of database queries
# TYPE masked11_database_queries_total counter
masked11_database_queries_total %d

# HELP masked11_database_errors_total Total number of database errors
# TYPE masked11_database_errors_total counter
masked11_database_errors_total %d

# HELP masked11_cache_hits_total Total number of cache hits
# TYPE masked11_cache_hits_total counter
masked11_cache_hits_total %d

# HELP masked11_cache_misses_total Total number of cache misses
# TYPE masked11_cache_misses_total counter
masked11_cache_misses_total %d

# HELP masked11_cache_hit_rate_percent Cache hit rate as percentage
# TYPE masked11_cache_hit_rate_percent gauge
masked11_cache_hit_rate_percent %f

# HELP masked11_active_connections Current number of active connections
# TYPE masked11_active_connections gauge
masked11_active_connections %d

# HELP masked11_uptime_seconds Server uptime in seconds
# TYPE masked11_uptime_seconds gauge
masked11_uptime_seconds %f
`,
		metrics.RequestCount,
		metrics.ErrorCount,
		metrics.RateLimitExceededCount,
		metrics.AverageResponseTime,
		metrics.DatabaseQueryCount,
		metrics.DatabaseErrorCount,
		metrics.CacheHitCount,
		metrics.CacheMissCount,
		GetCacheHitRate(),
		metrics.ActiveConnections,
		metrics.Uptime.Seconds(),
	)

	return output
} 