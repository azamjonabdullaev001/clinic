package models

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	FirstName  string    `gorm:"not null" json:"first_name"`
	LastName   string    `gorm:"not null" json:"last_name"`
	MiddleName string    `json:"middle_name"`
	Phone      string    `gorm:"uniqueIndex;not null" json:"phone"`
	Password   string    `gorm:"not null" json:"-"`
	CreatedAt  time.Time `json:"created_at"`
}

type Admin struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Phone    string `gorm:"uniqueIndex;not null" json:"phone"`
	Password string `gorm:"not null" json:"-"`
}

type Product struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"not null" json:"name"`
	Description     string    `json:"description"`
	QuantityPerPack int       `gorm:"not null" json:"quantity_per_pack"`
	PricePerPill    float64   `gorm:"not null" json:"price_per_pill"`
	PricePerPack    float64   `gorm:"-" json:"price_per_pack"`
	ImagePath       string    `json:"image_path"`
	CreatedAt       time.Time `json:"created_at"`
}

func (p *Product) ComputePackPrice() {
	p.PricePerPack = p.PricePerPill * float64(p.QuantityPerPack)
}

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `gorm:"not null" json:"user_id"`
	User      User        `json:"user"`
	Items     []OrderItem `json:"items"`
	Status    string      `gorm:"default:'pending';not null" json:"status"`
	Phone     string      `gorm:"not null" json:"phone"`
	CreatedAt time.Time   `json:"created_at"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	UnitType  string  `gorm:"default:'pack';not null" json:"unit_type"`
	Price     float64 `gorm:"not null" json:"price"`
}
