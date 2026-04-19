package database

import (
	"clinic-backend/config"
	"clinic-backend/models"
	"fmt"
	"log"

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
	log.Println("Database connected successfully")
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migrated successfully")
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
