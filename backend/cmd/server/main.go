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
	"github.com/Shrey-Yash/Masked11/internal/repositories/redis"
	"github.com/Shrey-Yash/Masked11/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}
	// Database
	database.InitMongo()

	// Redis
	database.InitRedis()
	sessionRepo := redisrepo.NewSessionRepository(database.Redis, database.Ctx)

	// Repo
	userRepo := mongodb.NewUserRepository(database.Mongo)
	productRepo := mongodb.NewProductRepository(database.Mongo)

	// Services
	authService := services.NewAuthService(userRepo)
	productService := services.NewProductService(productRepo)

	// Admin
	if err := authService.BootstrapAdmin(); err != nil {
		log.Println("Admin bootstrap failed:", err)
	}

	// Handlers
	authHandler := handlers.NewAuthHandler(authService, sessionRepo)
	userHandler := handlers.NewUserHandler(authService)
	productHandler := handlers.NewProductHandler(productService)

	app := fiber.New()

	// User auth
	app.Post("/api/register", authHandler.RegisterUser)
	app.Post("/api/login", authHandler.LoginUser)
	app.Post("/api/logout", authHandler.Logout)

	// User CRUD
	api := app.Group("/api")
	api.Use(middleware.JWTMiddleware())
	api.Get("/user", userHandler.GetUser)
	api.Put("/user", userHandler.UpdateUser)
	api.Delete("/user", userHandler.DeleteUser)

	// Product Routes (admin only)
	productGroup := app.Group("/api/products", middleware.AdminOnly())
	productGroup.Post("/", productHandler.CreateProduct)
	productGroup.Put("/:id", productHandler.UpdateProduct)
	productGroup.Delete("/:id", productHandler.DeleteProduct)

	// Public product routes
	app.Get("/api/products", productHandler.GetAllProducts)
	app.Get("/api/products/:id", productHandler.GetProductByID)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is running...")
	})	

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
