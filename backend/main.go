package main

import (
	"clinic-backend/config"
	"clinic-backend/database"
	"clinic-backend/handlers"
	"clinic-backend/middleware"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET environment variable must be set")
	}

	database.Connect(cfg)
	database.Migrate()
	database.SeedAdmin(cfg)

	r := gin.Default()
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatal("Failed to set trusted proxies:", err)
	}

	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			// Allow requests from the same host (HTTP and HTTPS) and from
			// localhost variants used during development.
			return origin == "" ||
				origin == "http://localhost:5173" ||
				origin == "http://localhost:3000" ||
				origin == "http://127.0.0.1:5173" ||
				origin == "http://127.0.0.1:3000"
		},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.Static("/uploads", "./uploads")
	r.MaxMultipartMemory = 10 << 20

	// Throttle each client IP. Generous enough for real panels (which poll and
	// open several requests per page) but enough to stop floods hitting Postgres.
	r.Use(middleware.RateLimit(30, 60))

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
				protected.POST("/payment-qr", handlers.UploadPaymentQR)
				protected.GET("/products", handlers.GetProducts)
				protected.POST("/products", handlers.CreateProduct)
				protected.PUT("/products/:id", handlers.UpdateProduct)
				protected.PUT("/products/:id/online", handlers.SetProductOnlineAvailability)
				protected.DELETE("/products/:id", handlers.DeleteProduct)
				protected.POST("/products/:id/image", handlers.UploadProductImage)
				protected.GET("/orders", handlers.GetOrders)
				protected.PUT("/orders/:id/status", handlers.UpdateOrderStatus)
				protected.DELETE("/orders/:id", handlers.DeleteOrder)
				protected.DELETE("/orders", handlers.DeleteAllOrders)
				protected.GET("/workers", handlers.GetWorkers)
				protected.POST("/workers", handlers.CreateWorker)
				protected.PUT("/workers/:id", handlers.UpdateWorker)
				protected.DELETE("/workers/:id", handlers.DeleteWorker)
				protected.GET("/workers/:id/stats", handlers.CacheGET(15*time.Second), handlers.GetWorkerStats)
				protected.GET("/faqs", handlers.GetFAQs)
				protected.POST("/faqs", handlers.CreateFAQ)
				protected.PUT("/faqs/:id", handlers.UpdateFAQ)
				protected.DELETE("/faqs/:id", handlers.DeleteFAQ)
				protected.GET("/support/threads", handlers.GetSupportThreads)
				protected.GET("/support/threads/:id", handlers.GetSupportThreadByID)
				protected.POST("/support/threads/:id/reply", handlers.ReplySupportThread)
				protected.GET("/analytics", handlers.CacheGET(15*time.Second), handlers.GetAnalytics)
				protected.GET("/products/:id/analytics", handlers.CacheGET(15*time.Second), handlers.GetProductAnalytics)
				protected.GET("/news", handlers.GetNewsPosts)
				protected.POST("/news", handlers.CreateNewsPost)
				protected.PUT("/news/:id", handlers.UpdateNewsPost)
				protected.DELETE("/news/:id", handlers.DeleteNewsPost)
				protected.DELETE("/news/images/:id", handlers.DeleteNewsImage)
				protected.DELETE("/product-comments/:commentId", handlers.DeleteProductCommentByID)
				protected.GET("/doctors", handlers.GetDoctors)
				protected.POST("/doctors", handlers.CreateDoctor)
				protected.PUT("/doctors/:id", handlers.UpdateDoctor)
				protected.DELETE("/doctors/:id", handlers.DeleteDoctor)
				protected.GET("/doctors/:id/stats", handlers.CacheGET(15*time.Second), handlers.GetDoctorStats)
				protected.GET("/marketologs", handlers.GetMarketologs)
				protected.GET("/marketologs/:id/stats", handlers.CacheGET(15*time.Second), handlers.GetMarketologStatsAdmin)
				protected.GET("/bts-branches", handlers.GetBtsBranches)
				protected.POST("/bts-branches", handlers.CreateBtsBranch)
				protected.PUT("/bts-branches/:id", handlers.UpdateBtsBranch)
				protected.DELETE("/bts-branches/:id", handlers.DeleteBtsBranch)
			}
		}

		products := api.Group("/products")
		{
			products.GET("", handlers.GetPublicProducts)
			products.GET("/:id", handlers.GetProduct)
			products.GET("/:id/comments", handlers.GetProductComments)
			products.POST("/:id/comments", middleware.UserAuth(), handlers.AddProductComment)
		}

		api.GET("/settings/payment-qr", handlers.GetPaymentQR)
		api.GET("/bts-branches", handlers.GetBtsBranches)

		api.GET("/faqs", handlers.GetFAQs)

		api.GET("/news", handlers.GetNewsPosts)

		api.GET("/doctors", handlers.GetDoctors)

		api.GET("/ws/stock", handlers.StockWebSocket)

		api.POST("/contact", handlers.SendContactMessage)

		orders := api.Group("/orders")
		orders.Use(middleware.UserAuth())
		{
			orders.POST("", handlers.CreateOrder)
			orders.GET("", handlers.GetUserOrders)
			orders.POST("/:id/receipt", handlers.UploadOrderReceipt)
			orders.POST("/:id/receipts", handlers.AddOrderReceipt) // split-payment: add another receipt photo
			orders.PUT("/:id/location", handlers.UpdateOrderLocation) // correct delivery address while still pending
			orders.DELETE("/:id/hide", handlers.HideUserOrder)
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
			pickup.POST("/orders/:id/return", handlers.ReturnOrderFull)
			pickup.DELETE("/orders/:id", handlers.DeletePickupOrder)
			pickup.PUT("/orders/:id/bts", handlers.UpdateBtsInfo)
			pickup.PUT("/orders/:id/notes", handlers.UpdateOrderNotes)
			pickup.PUT("/orders/:id/payment", handlers.UpdateOrderPayment)
			pickup.GET("/bts-branches", handlers.GetBtsBranches)
			pickup.POST("/bts-branches", handlers.CreateBtsBranch)
			pickup.PUT("/bts-branches/:id", handlers.UpdateBtsBranch)
			pickup.DELETE("/bts-branches/:id", handlers.DeleteBtsBranch)
			pickup.POST("/offline-sale", handlers.CreateOfflineSale)
			pickup.GET("/analytics", handlers.CacheGET(15*time.Second), handlers.GetWorkerAnalytics)
			pickup.GET("/stock", handlers.GetGlobalStock)
			pickup.POST("/stock", handlers.AddProductStock)
			// Product management from the pickup warehouse (same as the admin panel).
			// all-products returns the FULL catalogue (incl. online-hidden) for offline sales.
			pickup.GET("/all-products", handlers.GetProducts)
			pickup.POST("/products", handlers.CreateProduct)
			pickup.PUT("/products/:id", handlers.UpdateProduct)
			pickup.PUT("/products/:id/online", handlers.SetProductOnlineAvailability)
			pickup.DELETE("/products/:id", handlers.DeleteProduct)
			pickup.POST("/products/:id/image", handlers.UploadProductImage)
			pickup.GET("/marketologs", handlers.GetMarketologs)
			pickup.GET("/marketologs/:id/debt", handlers.GetMarketologDebt)
			pickup.POST("/marketologs/:id/payment", handlers.AddMarketologPayment)
			pickup.GET("/support/threads", handlers.GetWorkerSupportThreads)
			pickup.GET("/support/threads/:id", handlers.GetWorkerSupportThreadByID)
			pickup.POST("/support/threads/:id/reply", handlers.ReplyWorkerSupportThread)
		}

		manager := api.Group("/manager")
		manager.Use(middleware.WorkerAuth())
		{
			manager.GET("/products", handlers.GetProducts)
			manager.POST("/sale", handlers.CreateOfflineSale)
			manager.GET("/orders", handlers.GetManagerOrders)
			manager.GET("/analytics", handlers.CacheGET(15*time.Second), handlers.GetMarketologOwnAnalytics)
			manager.GET("/stock", handlers.GetGlobalStock)
			manager.POST("/stock", handlers.AddProductStock)
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
			doctor.GET("/analytics", handlers.CacheGET(15*time.Second), handlers.GetDoctorAnalytics)
		}
	}

	// Explicit http.Server so we can set timeouts and shut down gracefully.
	// WriteTimeout is intentionally left at 0 — the stock WebSocket holds a long-lived
	// connection that a write deadline would kill. ReadHeaderTimeout still guards against
	// slow-loris clients holding sockets open.
	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	go func() {
		log.Printf("Server starting on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Graceful shutdown: stop accepting new requests, let in-flight ones finish, then
	// close the DB pool. Prevents dropped requests and leaked connections on deploy.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}
	if sqlDB, err := database.DB.DB(); err == nil {
		_ = sqlDB.Close()
	}
	log.Println("Server stopped")
}
