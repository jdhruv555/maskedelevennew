package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds all configuration settings
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Security SecurityConfig
	Cache    CacheConfig
	Logging  LoggingConfig
	Monitoring MonitoringConfig
	Performance PerformanceConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port            string
	Environment     string
	Name            string
	Version         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	MongoDB MongoDBConfig
	PostgreSQL PostgreSQLConfig
	Redis   RedisConfig
}

// MongoDBConfig holds MongoDB-specific configuration
type MongoDBConfig struct {
	URI          string
	Database     string
	Username     string
	Password     string
	AuthSource   string
	MaxPoolSize  int
	MinPoolSize  int
	MaxConnIdleTime time.Duration
	MaxConnLifetime time.Duration
}

// PostgreSQLConfig holds PostgreSQL-specific configuration
type PostgreSQLConfig struct {
	URL             string
	Host            string
	Port            string
	Database        string
	Username        string
	Password        string
	SSLMode         string
	MaxConnections  int
	MinConnections  int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

// RedisConfig holds Redis-specific configuration
type RedisConfig struct {
	URI           string
	Password      string
	DB            int
	PoolSize      int
	MinIdleConns  int
	MaxRetries    int
	DialTimeout   time.Duration
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	IdleTimeout   time.Duration
	PoolTimeout   time.Duration
	MaxConnAge    time.Duration
}

// SecurityConfig holds security-related configuration
type SecurityConfig struct {
	SessionSecret    string
	JWTSecret        string
	JWTExpiry        time.Duration
	BCryptCost       int
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	RateLimitEnabled bool
	RateLimitRequests int
	RateLimitWindow  time.Duration
}

// CacheConfig holds cache-related configuration
type CacheConfig struct {
	Enabled  bool
	TTL      time.Duration
	MaxSize  int
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string
	Format string
	Output string
	File   string
}

// MonitoringConfig holds monitoring configuration
type MonitoringConfig struct {
	MetricsEnabled    bool
	MetricsPort       string
	HealthCheckEnabled bool
	ProfilingEnabled  bool
}

// PerformanceConfig holds performance-related configuration
type PerformanceConfig struct {
	GOMAXPROCS int
	GOGC       int
	GODEBUG    string
	CompressionEnabled bool
	CompressionLevel   int
	CompressionMinSize int
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server:   loadServerConfig(),
		Database: loadDatabaseConfig(),
		Security: loadSecurityConfig(),
		Cache:    loadCacheConfig(),
		Logging:  loadLoggingConfig(),
		Monitoring: loadMonitoringConfig(),
		Performance: loadPerformanceConfig(),
	}
}

func loadServerConfig() ServerConfig {
	return ServerConfig{
		Port:            getEnv("APP_PORT", "8080"),
		Environment:     getEnv("APP_ENV", "development"),
		Name:            getEnv("APP_NAME", "Masked11-API"),
		Version:         getEnv("APP_VERSION", "1.0.0"),
		ReadTimeout:     getDurationEnv("READ_TIMEOUT", 30*time.Second),
		WriteTimeout:    getDurationEnv("WRITE_TIMEOUT", 30*time.Second),
		IdleTimeout:     getDurationEnv("IDLE_TIMEOUT", 120*time.Second),
		ShutdownTimeout: getDurationEnv("SHUTDOWN_TIMEOUT", 30*time.Second),
	}
}

func loadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		MongoDB: MongoDBConfig{
			URI:             getEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database:        getEnv("MONGO_DB", "masked11"),
			Username:        getEnv("MONGO_USER", ""),
			Password:        getEnv("MONGO_PASSWORD", ""),
			AuthSource:      getEnv("MONGO_AUTH_SOURCE", "admin"),
			MaxPoolSize:     getIntEnv("MONGO_MAX_POOL_SIZE", 100),
			MinPoolSize:     getIntEnv("MONGO_MIN_POOL_SIZE", 5),
			MaxConnIdleTime: getDurationEnv("MONGO_MAX_CONN_IDLE_TIME", 30*time.Minute),
			MaxConnLifetime: getDurationEnv("MONGO_MAX_CONN_LIFETIME", 1*time.Hour),
		},
		PostgreSQL: PostgreSQLConfig{
			URL:             getEnv("POSTGRES_URL", "postgres://masked11_user:masked11_password@localhost:5432/masked11?sslmode=disable"),
			Host:            getEnv("POSTGRES_HOST", "localhost"),
			Port:            getEnv("POSTGRES_PORT", "5432"),
			Database:        getEnv("POSTGRES_DB", "masked11"),
			Username:        getEnv("POSTGRES_USER", "masked11_user"),
			Password:        getEnv("POSTGRES_PASSWORD", "masked11_password"),
			SSLMode:         getEnv("POSTGRES_SSL_MODE", "disable"),
			MaxConnections:  getIntEnv("POSTGRES_MAX_CONNECTIONS", 100),
			MinConnections:  getIntEnv("POSTGRES_MIN_CONNECTIONS", 5),
			ConnMaxLifetime: getDurationEnv("POSTGRES_CONN_MAX_LIFETIME", 1*time.Hour),
			ConnMaxIdleTime: getDurationEnv("POSTGRES_CONN_MAX_IDLE_TIME", 30*time.Minute),
		},
		Redis: RedisConfig{
			URI:           getEnv("REDIS_URI", "localhost:6379"),
			Password:      getEnv("REDIS_PASSWORD", ""),
			DB:            getIntEnv("REDIS_DB", 0),
			PoolSize:      getIntEnv("REDIS_POOL_SIZE", 50),
			MinIdleConns:  getIntEnv("REDIS_MIN_IDLE_CONNS", 10),
			MaxRetries:    getIntEnv("REDIS_MAX_RETRIES", 3),
			DialTimeout:   getDurationEnv("REDIS_DIAL_TIMEOUT", 5*time.Second),
			ReadTimeout:   getDurationEnv("REDIS_READ_TIMEOUT", 3*time.Second),
			WriteTimeout:  getDurationEnv("REDIS_WRITE_TIMEOUT", 3*time.Second),
			IdleTimeout:   getDurationEnv("REDIS_IDLE_TIMEOUT", 5*time.Minute),
			PoolTimeout:   getDurationEnv("REDIS_POOL_TIMEOUT", 4*time.Second),
			MaxConnAge:    getDurationEnv("REDIS_MAX_CONN_AGE", 30*time.Minute),
		},
	}
}

func loadSecurityConfig() SecurityConfig {
	return SecurityConfig{
		SessionSecret:    getEnv("SESSION_SECRET", "your-super-secret-key-change-in-production"),
		JWTSecret:        getEnv("JWT_SECRET", "your-jwt-secret-key-change-in-production"),
		JWTExpiry:        getDurationEnv("JWT_EXPIRY", 168*time.Hour),
		BCryptCost:       getIntEnv("BCRYPT_COST", 12),
		AllowedOrigins:   getStringSliceEnv("ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
		AllowedMethods:   getStringSliceEnv("ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		AllowedHeaders:   getStringSliceEnv("ALLOWED_HEADERS", []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}),
		RateLimitEnabled: getBoolEnv("RATE_LIMIT_ENABLED", true),
		RateLimitRequests: getIntEnv("RATE_LIMIT_REQUESTS", 100),
		RateLimitWindow:  getDurationEnv("RATE_LIMIT_WINDOW", 1*time.Minute),
	}
}

func loadCacheConfig() CacheConfig {
	return CacheConfig{
		Enabled: getBoolEnv("CACHE_ENABLED", true),
		TTL:     getDurationEnv("CACHE_TTL", 1*time.Hour),
		MaxSize: getIntEnv("CACHE_MAX_SIZE", 1000),
	}
}

func loadLoggingConfig() LoggingConfig {
	return LoggingConfig{
		Level:  getEnv("LOG_LEVEL", "info"),
		Format: getEnv("LOG_FORMAT", "json"),
		Output: getEnv("LOG_OUTPUT", "stdout"),
		File:   getEnv("LOG_FILE", ""),
	}
}

func loadMonitoringConfig() MonitoringConfig {
	return MonitoringConfig{
		MetricsEnabled:    getBoolEnv("METRICS_ENABLED", true),
		MetricsPort:       getEnv("METRICS_PORT", "9090"),
		HealthCheckEnabled: getBoolEnv("HEALTH_CHECK_ENABLED", true),
		ProfilingEnabled:  getBoolEnv("PROFILING_ENABLED", false),
	}
}

func loadPerformanceConfig() PerformanceConfig {
	return PerformanceConfig{
		GOMAXPROCS:        getIntEnv("GOMAXPROCS", 1),
		GOGC:              getIntEnv("GOGC", 100),
		GODEBUG:           getEnv("GODEBUG", "netdns=go"),
		CompressionEnabled: getBoolEnv("COMPRESSION_ENABLED", true),
		CompressionLevel:   getIntEnv("COMPRESSION_LEVEL", 6),
		CompressionMinSize: getIntEnv("COMPRESSION_MIN_SIZE", 1024),
	}
}

// Helper functions for environment variable parsing

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getStringSliceEnv(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		// Simple comma-separated parsing
		// In production, you might want more sophisticated parsing
		return []string{value}
	}
	return defaultValue
}

// IsDevelopment returns true if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.Server.Environment == "development"
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	return c.Server.Environment == "production"
}

// IsDebug returns true if debug mode is enabled
func (c *Config) IsDebug() bool {
	return getBoolEnv("DEBUG", c.IsDevelopment())
} 