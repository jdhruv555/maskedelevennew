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

	database.InitMongo()
	database.InitRedis()

	sessionRepo := redisrepo.NewSessionRepository(database.Redis, database.Ctx)

	userRepo := mongodb.NewUserRepository(database.Mongo)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService, sessionRepo)
	userHandler := handlers.NewUserHandler(authService)

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

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
