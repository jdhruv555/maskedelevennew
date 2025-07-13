package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/joho/godotenv"

	"github.com/Shrey-Yash/Masked11/internal/database"
	"github.com/Shrey-Yash/Masked11/internal/handlers"
	"github.com/Shrey-Yash/Masked11/internal/middleware"
	"github.com/Shrey-Yash/Masked11/internal/repositories/mongodb"
	"github.com/Shrey-Yash/Masked11/internal/repositories/postgres"
	redisrepo "github.com/Shrey-Yash/Masked11/internal/repositories/redis"
	"github.com/Shrey-Yash/Masked11/internal/services"
	"github.com/Shrey-Yash/Masked11/internal/utils"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	// Initialize databases with connection pooling
	if err := initializeDatabases(); err != nil {
		log.Fatal("Failed to initialize databases:", err)
	}

	// Run database migrations
	if err := scripts.MigrateDb(); err != nil {
		log.Printf("Migration warning: %v", err)
	}

	// Initialize repositories with connection pooling
	repos := initializeRepositories()

	// Initialize services
	services := initializeServices(repos)

	// Bootstrap admin user
	if err := services.AuthService.BootstrapAdmin(); err != nil {
		log.Println("Admin bootstrap failed:", err)
	}

	// Initialize handlers
	handlers := initializeHandlers(services)

	// Create Fiber app with optimized configuration
	app := createFiberApp()

	// Setup routes
	setupRoutes(app, handlers)

	// Start server with graceful shutdown
	startServer(app)
}

func initializeDatabases() error {
	// Initialize MongoDB with optimized settings
	if err := database.InitMongo(); err != nil {
		return err
	}

	// Initialize Redis with connection pooling
	if err := database.InitRedis(); err != nil {
		return err
	}

	// Initialize PostgreSQL with connection pooling
	if err := database.InitPostgres(); err != nil {
		return err
	}

	log.Println("✅ All databases initialized successfully")
	return nil
}

func initializeRepositories() map[string]interface{} {
	// Redis Repos with connection pooling
	sessionRepo := redisrepo.NewSessionRepository(database.Redis, database.Ctx)
	cartRepo := redisrepo.NewCartRepository(database.Redis, database.Ctx)

	// MongoDB Repos with optimized queries
	userRepo := mongodb.NewUserRepository(database.Mongo)
	productRepo := mongodb.NewProductRepository(database.Mongo)

	// PostgreSQL Repos with connection pooling
	orderRepo := postgres.NewOrderRepository(database.PostgresPool)

	return map[string]interface{}{
		"sessionRepo": sessionRepo,
		"cartRepo":    cartRepo,
		"userRepo":    userRepo,
		"productRepo": productRepo,
		"orderRepo":   orderRepo,
	}
}

func initializeServices(repos map[string]interface{}) map[string]interface{} {
	// Initialize services with dependency injection
	authService := services.NewAuthService(repos["userRepo"].(interfaces.UserRepository))
	productService := services.NewProductService(repos["productRepo"].(interfaces.ProductRepository))
	orderService := services.NewOrderService(
		repos["orderRepo"].(interfaces.OrderRepository),
		repos["cartRepo"].(interfaces.CartRepository),
		repos["userRepo"].(interfaces.UserRepository),
	)

	return map[string]interface{}{
		"AuthService":    authService,
		"ProductService": productService,
		"OrderService":   orderService,
	}
}

func initializeHandlers(services map[string]interface{}) map[string]interface{} {
	// Initialize handlers with services
	authHandler := handlers.NewAuthHandler(
		services["AuthService"].(*services.AuthService),
		repos["sessionRepo"].(interfaces.SessionRepository),
	)
	userHandler := handlers.NewUserHandler(services["AuthService"].(*services.AuthService))
	productHandler := handlers.NewProductHandler(services["ProductService"].(*services.ProductService))
	cartHandler := handlers.NewCartHandler(repos["cartRepo"].(interfaces.CartRepository))
	orderHandler := handlers.NewOrderHandler(services["OrderService"].(*services.OrderService))

	return map[string]interface{}{
		"authHandler":    authHandler,
		"userHandler":    userHandler,
		"productHandler": productHandler,
		"cartHandler":    cartHandler,
		"orderHandler":   orderHandler,
	}
}

func createFiberApp() *fiber.App {
	// Create Fiber app with optimized configuration
	app := fiber.New(fiber.Config{
		AppName:                  "Masked 11 Ecommerce API",
		ServerHeader:             "Masked11-API",
		DisableStartupMessage:    true,
		ReadTimeout:              30 * time.Second,
		WriteTimeout:             30 * time.Second,
		IdleTimeout:              120 * time.Second,
		BodyLimit:                10 * 1024 * 1024, // 10MB
		EnableTrustedProxyCheck:  true,
		ProxyHeader:              "X-Forwarded-For",
		EnableSplittingOnParsers: true,
		JSONEncoder:              utils.JSONEncoder,
		JSONDecoder:              utils.JSONDecoder,
	})

	// Security middleware
	app.Use(helmet.New(helmet.Config{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: false,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline';",
	}))

	// CORS middleware with optimized settings
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ALLOWED_ORIGINS"),
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Requested-With",
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	// Request ID middleware for tracing
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			return utils.GenerateRequestID()
		},
	}))

	// Structured logging middleware
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${method} | ${path} | ${ip} | ${requestID}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "UTC",
	}))

	// Recovery middleware
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Compression middleware
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Cache middleware for static responses
	app.Use(cache.New(cache.Config{
		Expiration:   1 * time.Hour,
		CacheControl: true,
	}))

	// Metrics middleware for tracking
	app.Use(utils.MetricsMiddleware())

	// Rate limiting middleware with Redis storage
	app.Use(middleware.RedisRateLimiter())

	// Timeout middleware for long-running requests
	app.Use(timeout.New(timeout.Config{
		Timeout: 30 * time.Second,
		Handler: func(c *fiber.Ctx) error {
			return fiber.NewError(fiber.StatusRequestTimeout, "Request timeout")
		},
	}))

	return app
}

func setupRoutes(app *fiber.App, handlers map[string]interface{}) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return utils.SuccessResponse(c, fiber.StatusOK, "Service is healthy", fiber.Map{
			"status":    "healthy",
			"timestamp": time.Now().UTC(),
			"version":   "1.0.0",
			"uptime":    utils.GetUptime().String(),
		})
	})

	// API status endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return utils.SuccessResponse(c, fiber.StatusOK, "Masked 11 Ecommerce API is running", fiber.Map{
			"message":   "Masked 11 Ecommerce API is running...",
			"version":   "1.0.0",
			"status":    "healthy",
			"timestamp": time.Now().UTC(),
			"uptime":    utils.GetUptime().String(),
		})
	})

	// API metrics endpoint
	app.Get("/metrics", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		metrics := utils.GetMetrics()
		return utils.SuccessResponse(c, fiber.StatusOK, "Metrics retrieved successfully", metrics)
	})

	// Prometheus metrics endpoint
	app.Get("/metrics/prometheus", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendString(utils.ExportMetricsForPrometheus())
	})

	// Database metrics endpoint
	app.Get("/metrics/database", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		dbMetrics := utils.GetDatabaseMetrics()
		return utils.SuccessResponse(c, fiber.StatusOK, "Database metrics retrieved successfully", dbMetrics)
	})

	// System metrics endpoint
	app.Get("/metrics/system", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		sysMetrics := utils.GetSystemMetrics()
		return utils.SuccessResponse(c, fiber.StatusOK, "System metrics retrieved successfully", sysMetrics)
	})

	// Authentication routes
	app.Post("/api/register", handlers["authHandler"].(*handlers.AuthHandler).RegisterUser)
	app.Post("/api/login", handlers["authHandler"].(*handlers.AuthHandler).LoginUser)
	app.Post("/api/logout", handlers["authHandler"].(*handlers.AuthHandler).Logout)

	// Public product routes with caching
	app.Get("/api/products", middleware.CacheMiddleware(), handlers["productHandler"].(*handlers.ProductHandler).GetAllProducts)
	app.Get("/api/products/search", handlers["productHandler"].(*handlers.ProductHandler).SearchProducts)
	app.Get("/api/products/categories", middleware.CacheMiddleware(), handlers["productHandler"].(*handlers.ProductHandler).GetProductCategories)
	app.Get("/api/products/categories/:category", middleware.CacheMiddleware(), handlers["productHandler"].(*handlers.ProductHandler).GetProductsByCategory)
	app.Get("/api/products/featured", middleware.CacheMiddleware(), handlers["productHandler"].(*handlers.ProductHandler).GetFeaturedProducts)
	app.Get("/api/products/:id", middleware.CacheMiddleware(), handlers["productHandler"].(*handlers.ProductHandler).GetProductByID)

	// Protected user routes
	api := app.Group("/api", middleware.JWTMiddleware())
	api.Get("/user", handlers["userHandler"].(*handlers.UserHandler).GetUser)
	api.Put("/user", handlers["userHandler"].(*handlers.UserHandler).UpdateUser)
	api.Delete("/user", handlers["userHandler"].(*handlers.UserHandler).DeleteUser)

	// Product management (admin only)
	productGroup := app.Group("/api/admin/products", middleware.AdminOnly())
	productGroup.Post("/", handlers["productHandler"].(*handlers.ProductHandler).CreateProduct)
	productGroup.Put("/:id", handlers["productHandler"].(*handlers.ProductHandler).UpdateProduct)
	productGroup.Delete("/:id", handlers["productHandler"].(*handlers.ProductHandler).DeleteProduct)

	// Cart routes
	app.Post("/api/cart/add", handlers["cartHandler"].(*handlers.CartHandler).AddToCart)
	app.Get("/api/cart", handlers["cartHandler"].(*handlers.CartHandler).GetCart)
	app.Delete("/api/cart/remove/:id", handlers["cartHandler"].(*handlers.CartHandler).RemoveFromCart)
	app.Delete("/api/cart/clear", handlers["cartHandler"].(*handlers.CartHandler).ClearCart)

	// Order routes
	orderGroup := app.Group("/api/orders", middleware.JWTMiddleware())
	orderGroup.Post("/", handlers["orderHandler"].(*handlers.OrderHandler).CreateOrder)
	orderGroup.Get("/", handlers["orderHandler"].(*handlers.OrderHandler).GetOrdersByUserID)
	orderGroup.Get("/:id", handlers["orderHandler"].(*handlers.OrderHandler).GetOrderByID)
	orderGroup.Put("/:id/cancel", handlers["orderHandler"].(*handlers.OrderHandler).CancelOrder)

	adminOrderGroup := app.Group("/api/admin/orders", middleware.AdminOnly())
	adminOrderGroup.Put("/:id/status", handlers["orderHandler"].(*handlers.OrderHandler).UpdateOrderStatus)
	adminOrderGroup.Delete("/:id", handlers["orderHandler"].(*handlers.OrderHandler).DeleteOrder)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return utils.NotFoundResponse(c, "The requested resource was not found")
	})
}

func startServer(app *fiber.App) {
	// Get port from environment
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Create channel for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		log.Printf("🚀 Server starting on port %s", port)
		log.Printf("📊 Health check: http://localhost:%s/health", port)
		log.Printf("📈 Metrics: http://localhost:%s/metrics", port)

		if err := app.Listen(":" + port); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Wait for interrupt signal
	<-quit
	log.Println("🛑 Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("✅ Server exited gracefully")
}

var startTime = time.Now()
