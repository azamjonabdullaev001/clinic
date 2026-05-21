package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"not null" json:"first_name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	MiddleName string   `json:"middle_name"`
	Phone     string    `gorm:"uniqueIndex;not null" json:"phone"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Doctor struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Specialty string    `json:"specialty"`
	Phone     string    `gorm:"uniqueIndex" json:"phone"`
	Password  string    `gorm:"default:''" json:"-"`
	CreatedAt time.Time `json:"created_at"`
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
	Role     string `gorm:"default:'pickup';not null" json:"role"` // "pickup" or "nurse"
}

type Product struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"not null" json:"name"`
	Description     string    `json:"description"`
	QuantityPerPack int       `gorm:"not null" json:"quantity_per_pack"`
	PricePerPill    float64   `gorm:"not null" json:"price_per_pill"`
	PricePerPack    float64   `gorm:"-" json:"price_per_pack"`
	ImagePath       string    `json:"image_path"`
	StockQuantity   int       `gorm:"default:0" json:"stock_quantity"`
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

type NewsPost struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Title       string      `gorm:"not null" json:"title"`
	Description string      `gorm:"type:text" json:"description"`
	ImagePath   string      `json:"image_path"`
	VideoURL    string      `json:"video_url"`
	Images      []NewsImage `gorm:"foreignKey:NewsPostID;constraint:OnDelete:CASCADE" json:"images"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type NewsImage struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	NewsPostID uint      `gorm:"index;not null" json:"news_post_id"`
	ImagePath  string    `gorm:"not null" json:"image_path"`
	Position   int       `gorm:"default:0" json:"position"`
	CreatedAt  time.Time `json:"created_at"`
}

func (p *Product) ComputePackPrice() {
	p.PricePerPack = p.PricePerPill * float64(p.QuantityPerPack)
}

type Order struct {
	ID              uint        `gorm:"primaryKey" json:"id"`
	UserID          *uint       `json:"user_id"`
	User            User        `json:"user"`
	WorkerID        *uint       `json:"worker_id"`
	DoctorID        *uint       `json:"doctor_id"`
	Items           []OrderItem `json:"items"`
	Status          string      `gorm:"default:'pending';not null" json:"status"`
	Phone           string      `gorm:"not null" json:"phone"`
	OrderCode       string      `gorm:"uniqueIndex;size:10" json:"order_code"`
	DeliveryAddress string      `json:"delivery_address"`
	Latitude        float64     `json:"latitude"`
	Longitude       float64     `json:"longitude"`
	ReferredBy      string      `json:"referred_by"`
	IsOffline       bool        `gorm:"default:false" json:"is_offline"`
	IsNurseOrder    bool        `gorm:"default:false" json:"is_nurse_order"`
	OfflineNote     string      `json:"offline_note"`
	PatientFName    string      `json:"patient_first_name"`
	PatientLName    string      `json:"patient_last_name"`
	CreatedAt       time.Time   `json:"created_at"`
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
