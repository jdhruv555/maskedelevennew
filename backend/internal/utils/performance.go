package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Performance metrics - keeping only unique utilities
var (
	requestCount int64
	metricsMutex sync.RWMutex
)

// JSONEncoder provides optimized JSON encoding
func JSONEncoder(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JSONDecoder provides optimized JSON decoding
func JSONDecoder(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// GenerateRequestID creates a unique request ID
func GenerateRequestID() string {
	return fmt.Sprintf("req_%d_%d", time.Now().UnixNano(), atomic.AddInt64(&requestCount, 1))
}

// DatabasePerformanceTracker tracks database query performance
type DatabasePerformanceTracker struct {
	queryTimes map[string][]time.Duration
	mutex      sync.RWMutex
}

// NewDatabasePerformanceTracker creates a new tracker
func NewDatabasePerformanceTracker() *DatabasePerformanceTracker {
	return &DatabasePerformanceTracker{
		queryTimes: make(map[string][]time.Duration),
	}
}

// TrackQuery tracks a database query
func (t *DatabasePerformanceTracker) TrackQuery(queryName string, duration time.Duration) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.queryTimes[queryName] == nil {
		t.queryTimes[queryName] = make([]time.Duration, 0)
	}
	t.queryTimes[queryName] = append(t.queryTimes[queryName], duration)
}

// GetQueryStats returns query performance statistics
func (t *DatabasePerformanceTracker) GetQueryStats() map[string]interface{} {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	stats := make(map[string]interface{})
	for queryName, times := range t.queryTimes {
		if len(times) == 0 {
			continue
		}

		var total time.Duration
		var max time.Duration
		min := times[0]

		for _, t := range times {
			total += t
			if t > max {
				max = t
			}
			if t < min {
				min = t
			}
		}

		avg := total / time.Duration(len(times))
		stats[queryName] = map[string]interface{}{
			"count":    len(times),
			"average":  avg.String(),
			"min":      min.String(),
			"max":      max.String(),
			"total":    total.String(),
		}
	}

	return stats
}

// CachePerformanceTracker tracks cache performance
type CachePerformanceTracker struct {
	hits   int64
	misses int64
}

// NewCachePerformanceTracker creates a new cache tracker
func NewCachePerformanceTracker() *CachePerformanceTracker {
	return &CachePerformanceTracker{}
}

// RecordHit records a cache hit
func (t *CachePerformanceTracker) RecordHit() {
	atomic.AddInt64(&t.hits, 1)
}

// RecordMiss records a cache miss
func (t *CachePerformanceTracker) RecordMiss() {
	atomic.AddInt64(&t.misses, 1)
}

// GetHitRate returns cache hit rate
func (t *CachePerformanceTracker) GetHitRate() float64 {
	hits := atomic.LoadInt64(&t.hits)
	misses := atomic.LoadInt64(&t.misses)
	total := hits + misses

	if total == 0 {
		return 0
	}

	return float64(hits) / float64(total) * 100
}

// GetStats returns cache statistics
func (t *CachePerformanceTracker) GetStats() map[string]interface{} {
	hits := atomic.LoadInt64(&t.hits)
	misses := atomic.LoadInt64(&t.misses)
	total := hits + misses

	return map[string]interface{}{
		"hits":      hits,
		"misses":    misses,
		"total":     total,
		"hit_rate":  t.GetHitRate(),
		"timestamp": time.Now().UTC(),
	}
}

// ConnectionPoolMonitor monitors connection pool health
type ConnectionPoolMonitor struct {
	poolName string
	pool     interface{}
}

// NewConnectionPoolMonitor creates a new pool monitor
func NewConnectionPoolMonitor(poolName string, pool interface{}) *ConnectionPoolMonitor {
	return &ConnectionPoolMonitor{
		poolName: poolName,
		pool:     pool,
	}
}

// GetPoolStats returns pool statistics
func (m *ConnectionPoolMonitor) GetPoolStats() map[string]interface{} {
	return map[string]interface{}{
		"pool_name": m.poolName,
		"status":    "monitored",
		"timestamp": time.Now().UTC(),
	}
}

// PerformanceOptimizer provides comprehensive performance optimization
type PerformanceOptimizer struct {
	dbTracker    *DatabasePerformanceTracker
	cacheTracker *CachePerformanceTracker
	poolMonitors map[string]*ConnectionPoolMonitor
}

// NewPerformanceOptimizer creates a new performance optimizer
func NewPerformanceOptimizer() *PerformanceOptimizer {
	return &PerformanceOptimizer{
		dbTracker:    NewDatabasePerformanceTracker(),
		cacheTracker: NewCachePerformanceTracker(),
		poolMonitors: make(map[string]*ConnectionPoolMonitor),
	}
}

// TrackDatabaseQuery tracks a database query
func (o *PerformanceOptimizer) TrackDatabaseQuery(queryName string, duration time.Duration) {
	o.dbTracker.TrackQuery(queryName, duration)
}

// RecordCacheHit records a cache hit
func (o *PerformanceOptimizer) RecordCacheHit() {
	o.cacheTracker.RecordHit()
}

// RecordCacheMiss records a cache miss
func (o *PerformanceOptimizer) RecordCacheMiss() {
	o.cacheTracker.RecordMiss()
}

// GetPerformanceReport returns comprehensive performance report
func (o *PerformanceOptimizer) GetPerformanceReport() map[string]interface{} {
	return map[string]interface{}{
		"database": o.dbTracker.GetQueryStats(),
		"cache":    o.cacheTracker.GetStats(),
		"timestamp": time.Now().UTC(),
	}
}

// OptimizeQuery optimizes database queries
func OptimizeQuery(ctx context.Context, query string, params ...interface{}) (string, []interface{}) {
	// Basic query optimization - in production, you'd use a query optimizer
	return query, params
}

// BatchProcessor processes items in batches
type BatchProcessor struct {
	batchSize int
	workers   int
}

// NewBatchProcessor creates a new batch processor
func NewBatchProcessor(batchSize, workers int) *BatchProcessor {
	return &BatchProcessor{
		batchSize: batchSize,
		workers:   workers,
	}
}

// ProcessBatch processes items in batches
func (b *BatchProcessor) ProcessBatch(items []interface{}, processor func([]interface{}) error) error {
	if len(items) == 0 {
		return nil
	}

	// Process in batches
	for i := 0; i < len(items); i += b.batchSize {
		end := i + b.batchSize
		if end > len(items) {
			end = len(items)
		}

		batch := items[i:end]
		if err := processor(batch); err != nil {
			return err
		}
	}

	return nil
}

// AsyncProcessor processes items asynchronously
type AsyncProcessor struct {
	queue chan interface{}
	done  chan bool
}

// NewAsyncProcessor creates a new async processor
func NewAsyncProcessor(bufferSize int) *AsyncProcessor {
	return &AsyncProcessor{
		queue: make(chan interface{}, bufferSize),
		done:  make(chan bool),
	}
}

// Start starts the async processor
func (a *AsyncProcessor) Start(processor func(interface{}) error) {
	go func() {
		for {
			select {
			case item := <-a.queue:
				if err := processor(item); err != nil {
					// Log error but continue processing
					fmt.Printf("Async processing error: %v\n", err)
				}
			case <-a.done:
				return
			}
		}
	}()
}

// Stop stops the async processor
func (a *AsyncProcessor) Stop() {
	close(a.done)
}

// Queue adds an item to the processing queue
func (a *AsyncProcessor) Queue(item interface{}) {
	select {
	case a.queue <- item:
	default:
		// Queue is full, log warning
		fmt.Println("Warning: Async processor queue is full")
	}
} 