package main

import (
	"clinic-backend/config"
	"clinic-backend/database"
	"clinic-backend/handlers"
	"clinic-backend/middleware"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	database.Connect(cfg)
	database.Migrate()
	database.SeedAdmin(cfg)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.Static("/uploads", "./uploads")
	r.MaxMultipartMemory = 10 << 20

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		admin := api.Group("/admin")
		{
			admin.POST("/login", handlers.AdminLogin)

			protected := admin.Group("")
			protected.Use(middleware.AdminAuth())
			{
				protected.GET("/profile", handlers.AdminProfile)
				protected.PUT("/settings", handlers.UpdateAdminSettings)
				protected.GET("/products", handlers.GetProducts)
				protected.POST("/products", handlers.CreateProduct)
				protected.PUT("/products/:id", handlers.UpdateProduct)
				protected.DELETE("/products/:id", handlers.DeleteProduct)
				protected.POST("/products/:id/image", handlers.UploadProductImage)
				protected.GET("/orders", handlers.GetOrders)
				protected.PUT("/orders/:id/status", handlers.UpdateOrderStatus)
			}
		}

		products := api.Group("/products")
		{
			products.GET("", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
		}

		orders := api.Group("/orders")
		orders.Use(middleware.UserAuth())
		{
			orders.POST("", handlers.CreateOrder)
			orders.GET("", handlers.GetUserOrders)
		}
	}

	log.Printf("Server starting on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
