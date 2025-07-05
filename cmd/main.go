package main

import (
	"log"
	"net/http"

	"rest-api/internal/config"
	"rest-api/internal/handlers"
	"rest-api/internal/middleware"
	"rest-api/internal/repositories"
	"rest-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize validator
	validate := validator.New()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	// orderRepo := repositories.NewOrderRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo, validate)
	productService := services.NewProductService(productRepo, validate)
	// orderService := services.NewOrderService(orderRepo, productRepo, validate)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	// orderHandler := handlers.NewOrderHandler(orderService)

	// Setup router
	r := gin.Default()

	// Add middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ErrorHandler())

	// Add JWT config to context
	r.Use(func(c *gin.Context) {
		c.Set("jwt_secret", cfg.JWTSecret)
		c.Set("jwt_expire", cfg.JWTExpire)
		c.Next()
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.CreateUser)
			auth.POST("/login", userHandler.Login)
		}

		// User routes
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			users.GET("/profile", userHandler.GetProfile)
			users.PUT("/profile", userHandler.UpdateProfile)
			users.DELETE("/profile", userHandler.DeleteUser)
			users.GET("/", userHandler.GetAllUsers)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.GET("/", productHandler.GetAllProducts)
			products.GET("/:id", productHandler.GetProduct)

			// Protected product routes
			products.Use(middleware.AuthMiddleware(cfg.JWTSecret))
			products.POST("/", productHandler.CreateProduct)
			products.GET("/my", productHandler.GetMyProducts)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}

		// Order routes (if implemented)
		// orders := api.Group("/orders")
		// orders.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		// {
		// 	orders.POST("/", orderHandler.CreateOrder)
		// 	orders.GET("/", orderHandler.GetMyOrders)
		// 	orders.GET("/:id", orderHandler.GetOrder)
		// 	orders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		// }
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
