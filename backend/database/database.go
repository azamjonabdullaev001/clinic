package database

import (
	"clinic-backend/config"
	"clinic-backend/models"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Bound the connection pool. Without this the default is an UNLIMITED number of
	// open connections, which exhausts Postgres' max_connections under load and adds
	// connection-setup latency to every request.
	if sqlDB, e := DB.DB(); e == nil {
		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	log.Println("Database connected successfully")
}

func Migrate() {
	// Make orders.user_id nullable to support offline sales (no registered user)
	DB.Exec("ALTER TABLE orders ALTER COLUMN user_id DROP NOT NULL")
	// Set default role for existing workers
	DB.Exec("UPDATE workers SET role = 'pickup' WHERE role IS NULL OR role = ''")

	// Detect whether order_items.original_quantity already exists (for one-time backfill below)
	hadOrigQty := DB.Migrator().HasColumn(&models.OrderItem{}, "OriginalQuantity")
	// Detect whether products.stock_converted exists (one-time capsules->pieces conversion)
	hadStockConverted := DB.Migrator().HasColumn(&models.Product{}, "StockConverted")

	err := DB.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Worker{},
		&models.Product{},
		&models.FAQ{},
		&models.FAQAnswer{},
		&models.SupportThread{},
		&models.SupportMessage{},
		&models.Order{},
		&models.OrderItem{},
		&models.NewsPost{},
		&models.NewsImage{},
		&models.Doctor{},
		&models.WorkerStock{},
		&models.ProductComment{},
		&models.MarketologPayment{},
		&models.Setting{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// One-time backfill: existing items had no original quantity, so treat current
	// quantity as the originally ordered amount. Runs only when the column is first added.
	if !hadOrigQty {
		DB.Exec("UPDATE order_items SET original_quantity = quantity")
	}

	// One-time conversion: stock used to be counted in capsules; switch to pieces.
	if !hadStockConverted {
		DB.Exec("UPDATE products SET stock_quantity = stock_quantity * quantity_per_pack")
		DB.Exec("UPDATE products SET stock_converted = true")
	}

	DB.Exec("UPDATE workers SET role = 'pickup' WHERE role IS NULL OR role = ''")

	createPerformanceIndexes()
	log.Println("Database migrated successfully")
}

// createPerformanceIndexes adds indexes on the columns the hot query paths filter and
// join on. GORM's `not null` tags do NOT create indexes, so without these every analytics
// query, order-list page and Preload does a sequential scan of orders / order_items.
// All statements are idempotent (IF NOT EXISTS) and safe to run on every boot.
func createPerformanceIndexes() {
	stmts := []string{
		// Preloads (Order -> Items) and the product-analytics JOIN hit these constantly.
		"CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items (order_id)",
		"CREATE INDEX IF NOT EXISTS idx_order_items_product_id ON order_items (product_id)",
		// Analytics / order lists filter by time, status and the owning channel.
		"CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders (created_at)",
		"CREATE INDEX IF NOT EXISTS idx_orders_status ON orders (status)",
		"CREATE INDEX IF NOT EXISTS idx_orders_worker_id ON orders (worker_id)",
		"CREATE INDEX IF NOT EXISTS idx_orders_marketolog_id ON orders (marketolog_id)",
		"CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders (user_id)",
		// Composite indexes for the exact predicates used by GetAnalytics / WorkerAnalytics
		// / WorkerDashboard (status + time range, and per-worker reports).
		"CREATE INDEX IF NOT EXISTS idx_orders_status_created ON orders (status, created_at)",
		"CREATE INDEX IF NOT EXISTS idx_orders_worker_status_created ON orders (worker_id, status, created_at)",
	}
	for _, s := range stmts {
		if err := DB.Exec(s).Error; err != nil {
			log.Printf("index creation skipped (%s): %v", s, err)
		}
	}
}

func SeedAdmin(cfg *config.Config) {
	var count int64
	DB.Model(&models.Admin{}).Count(&count)
	if count > 0 {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash admin password:", err)
	}

	admin := models.Admin{
		Phone:    cfg.AdminPhone,
		Password: string(hash),
	}
	if err := DB.Create(&admin).Error; err != nil {
		log.Fatal("Failed to seed admin:", err)
	}
	log.Println("Admin user seeded successfully")
}
