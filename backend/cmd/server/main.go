package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/Shrey-Yash/Masked11/internal/database"
	"github.com/Shrey-Yash/Masked11/internal/handlers"
	"github.com/Shrey-Yash/Masked11/internal/middleware"
	"github.com/Shrey-Yash/Masked11/internal/repositories/mongodb"
	"github.com/Shrey-Yash/Masked11/internal/repositories/postgres"
	"github.com/Shrey-Yash/Masked11/internal/repositories/redis"
	"github.com/Shrey-Yash/Masked11/internal/services"
	"github.com/Shrey-Yash/Masked11/scripts"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	// Initialize databases
	database.InitMongo()
	database.InitRedis()
	database.InitPostgres()
	
	// Postgres Migration
	scripts.MigrateDb()

	// Redis Repos
	sessionRepo := redisrepo.NewSessionRepository(database.Redis, database.Ctx)
	cartRepo := redisrepo.NewCartRepository(database.Redis, database.Ctx)

	// MongoDB Repos
	userRepo := mongodb.NewUserRepository(database.Mongo)
	productRepo := mongodb.NewProductRepository(database.Mongo)

	// Postgres Repos
	orderRepo := postgres.NewOrderRepository(database.PostgresPool)

	// Services
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo, cartRepo, userRepo)

	// Admin Bootstrap
	if err := authService.BootstrapAdmin(); err != nil {
		log.Println("Admin bootstrap failed:", err)
	}

	// Handlers
	authHandler := handlers.NewAuthHandler(authService, sessionRepo)
	userHandler := handlers.NewUserHandler(authService)
	productHandler := handlers.NewProductHandler(productService)
	cartHandler := handlers.NewCartHandler(cartRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	app := fiber.New()

	// Public base route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running...")
	})

	// Authentication
	app.Post("/api/register", authHandler.RegisterUser)
	app.Post("/api/login", authHandler.LoginUser)
	app.Post("/api/logout", authHandler.Logout)

	// Protected user routes
	api := app.Group("/api")
	api.Use(middleware.JWTMiddleware())

	// User CRUD
	api.Get("/user", userHandler.GetUser)
	api.Put("/user", userHandler.UpdateUser)
	api.Delete("/user", userHandler.DeleteUser)

	// Product management (admin only)
	productGroup := app.Group("/api/admin/products", middleware.AdminOnly())
	productGroup.Post("/", productHandler.CreateProduct)
	productGroup.Put("/:id", productHandler.UpdateProduct)
	productGroup.Delete("/:id", productHandler.DeleteProduct)

	// Public product access
	app.Get("/api/products", productHandler.GetAllProducts)
	app.Get("/api/products/:id", productHandler.GetProductByID)

	// Cart routes
	app.Post("/api/cart/add", cartHandler.AddToCart)
	app.Get("/api/cart", cartHandler.GetCart)
	app.Delete("/api/cart/remove/:id", cartHandler.RemoveFromCart)
	app.Delete("/api/cart/clear", cartHandler.ClearCart)

	// Order routes
	orderGroup := app.Group("/api/orders", middleware.JWTMiddleware())
	orderGroup.Post("/", orderHandler.CreateOrder)
	orderGroup.Get("/", orderHandler.GetOrdersByUserID)
	orderGroup.Get("/:id", orderHandler.GetOrderByID)
	orderGroup.Put("/:id/cancel", orderHandler.CancelOrder)
	
	adminOrderGroup := app.Group("/api/admin/orders", middleware.AdminOnly())
	adminOrderGroup.Put("/:id/status", orderHandler.UpdateOrderStatus)
	adminOrderGroup.Delete("/:id", orderHandler.DeleteOrder)

	// Start server
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
