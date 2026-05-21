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
				protected.GET("/workers", handlers.GetWorkers)
				protected.POST("/workers", handlers.CreateWorker)
				protected.PUT("/workers/:id", handlers.UpdateWorker)
				protected.DELETE("/workers/:id", handlers.DeleteWorker)
				protected.GET("/faqs", handlers.GetFAQs)
				protected.POST("/faqs", handlers.CreateFAQ)
				protected.PUT("/faqs/:id", handlers.UpdateFAQ)
				protected.DELETE("/faqs/:id", handlers.DeleteFAQ)
				protected.GET("/support/threads", handlers.GetSupportThreads)
				protected.GET("/support/threads/:id", handlers.GetSupportThreadByID)
				protected.POST("/support/threads/:id/reply", handlers.ReplySupportThread)
				protected.GET("/analytics", handlers.GetAnalytics)
				protected.POST("/offline-sale", handlers.CreateOfflineSale)
				protected.GET("/news", handlers.GetNewsPosts)
				protected.POST("/news", handlers.CreateNewsPost)
				protected.PUT("/news/:id", handlers.UpdateNewsPost)
				protected.DELETE("/news/:id", handlers.DeleteNewsPost)
				protected.DELETE("/news/images/:id", handlers.DeleteNewsImage)
				protected.DELETE("/products/:id/comment", handlers.DeleteProductComment)
				protected.GET("/doctors", handlers.GetDoctors)
				protected.POST("/doctors", handlers.CreateDoctor)
				protected.PUT("/doctors/:id", handlers.UpdateDoctor)
				protected.DELETE("/doctors/:id", handlers.DeleteDoctor)
				protected.GET("/doctors/:id/stats", handlers.GetDoctorStats)
			}
		}

		products := api.Group("/products")
		{
			products.GET("", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
		}

		api.GET("/faqs", handlers.GetFAQs)

		api.GET("/news", handlers.GetNewsPosts)

		api.GET("/doctors", handlers.GetDoctors)

		api.POST("/contact", handlers.SendContactMessage)

		orders := api.Group("/orders")
		orders.Use(middleware.UserAuth())
		{
			orders.POST("", handlers.CreateOrder)
			orders.GET("", handlers.GetUserOrders)
		}

		support := api.Group("/support")
		support.Use(middleware.UserAuth())
		{
			support.GET("/thread", handlers.GetUserSupportThread)
			support.POST("/messages", handlers.SendUserSupportMessage)
			support.GET("/unread-count", handlers.GetUserUnreadCount)
		}

		pickup := api.Group("/pickup")
		pickup.Use(middleware.WorkerAuth())
		{
			pickup.GET("/orders", handlers.GetPickupOrders)
			pickup.GET("/orders/code/:code", handlers.GetOrderByCode)
			pickup.GET("/nurse-order/:code", handlers.ConfirmNurseOrder)
			pickup.PUT("/orders/:id/status", handlers.UpdatePickupOrderStatus)
			pickup.PUT("/orders/:id/items", handlers.UpdateOrderItems)
			pickup.POST("/offline-sale", handlers.CreateOfflineSale)
			pickup.GET("/support/threads", handlers.GetWorkerSupportThreads)
			pickup.GET("/support/threads/:id", handlers.GetWorkerSupportThreadByID)
			pickup.POST("/support/threads/:id/reply", handlers.ReplyWorkerSupportThread)
		}

		nurse := api.Group("/nurse")
		nurse.Use(middleware.WorkerAuth())
		{
			nurse.POST("/orders", handlers.CreateNurseOrder)
			nurse.GET("/orders", handlers.GetNurseOrders)
			nurse.GET("/products", handlers.GetProducts)
		}

		doctor := api.Group("/doctor")
		doctor.Use(middleware.DoctorAuth())
		{
			doctor.GET("/profile", handlers.GetDoctorProfile)
			doctor.POST("/orders", handlers.CreateDoctorOrder)
			doctor.GET("/orders", handlers.GetDoctorOrders)
			doctor.GET("/products", handlers.GetProducts)
			doctor.GET("/analytics", handlers.GetDoctorAnalytics)
		}
	}

	log.Printf("Server starting on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
