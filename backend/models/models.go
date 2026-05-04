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

type Worker struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
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

type FAQ struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Question  string      `gorm:"not null" json:"question"`
	Answers   []FAQAnswer `gorm:"constraint:OnDelete:CASCADE" json:"answers"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type FAQAnswer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FAQID     uint      `gorm:"index;not null" json:"faq_id"`
	Text      string    `gorm:"not null" json:"text"`
	Position  int       `gorm:"not null;default:0" json:"position"`
	CreatedAt time.Time `json:"created_at"`
}

type SupportThread struct {
	ID        uint             `gorm:"primaryKey" json:"id"`
	UserID    uint             `gorm:"uniqueIndex;not null" json:"user_id"`
	User      User             `gorm:"foreignKey:UserID" json:"user"`
	Messages  []SupportMessage `gorm:"foreignKey:ThreadID;constraint:OnDelete:CASCADE" json:"messages"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type SupportMessage struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ThreadID    uint      `gorm:"index;not null" json:"thread_id"`
	SenderRole  string    `gorm:"not null" json:"sender_role"`
	Message     string    `gorm:"type:text;not null" json:"message"`
	ReadByUser  bool      `gorm:"default:false" json:"read_by_user"`
	ReadByAdmin bool      `gorm:"default:false" json:"read_by_admin"`
	CreatedAt   time.Time `json:"created_at"`
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
	OrderCode string      `gorm:"uniqueIndex;size:6" json:"order_code"`
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
