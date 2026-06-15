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
	Role     string `gorm:"default:'pickup';not null" json:"role"` // "pickup", "nurse" or "manager"
}

type Product struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"not null" json:"name"`
	Description     string    `json:"description"`
	Category        string    `gorm:"default:''" json:"category"`
	QuantityPerPack int       `gorm:"not null" json:"quantity_per_pack"`
	PricePerPill    float64   `gorm:"not null" json:"price_per_pill"`
	PricePerPack    float64   `gorm:"-" json:"price_per_pack"`
	ImagePath       string    `json:"image_path"`
	StockQuantity   int       `gorm:"default:0" json:"stock_quantity"` // total stock in PIECES
	StockConverted  bool      `gorm:"default:false" json:"-"`          // migration marker: capsules -> pieces
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
	MarketologID    *uint       `json:"marketolog_id"` // marketolog this debt-sale belongs to
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
	IsVIP           bool        `gorm:"default:false" json:"is_vip"` // own hospital patient — free
	PaymentMethod   string      `json:"payment_method"`              // "cash", "terminal", "card", "online", "marketplace", "mixed"
	CardType        string      `json:"card_type"`                   // "cassa1", "click", "transfer" (legacy: humo/uzcard/visa/mastercard)
	PaymentSplits   string      `gorm:"type:text" json:"payment_splits"` // JSON [{"method":"cash","amount":2500000},...] when paid via several methods
	DiscountPercent float64     `gorm:"default:0" json:"discount_percent"` // cashier-applied discount on offline sales
	SalesChannel    string      `json:"sales_channel"`               // marketplace for manager sales: "ozon", "yandex", ...
	OfflineNote     string      `json:"offline_note"`
	PatientFName    string      `json:"patient_first_name"`
	PatientLName    string      `json:"patient_last_name"`
	CancellationReason string   `json:"cancellation_reason"`
	CancelledByName    string   `json:"cancelled_by_name"`
	CancelledByRole    string   `json:"cancelled_by_role"`
	IsReturned         bool     `gorm:"default:false" json:"is_returned"` // delivered order edited as a return
	ReturnReason       string   `json:"return_reason"`
	IsEdited           bool     `gorm:"default:false;index" json:"is_edited"` // items edited at pickup — shown on admin with red border
	IsDeleted          bool     `gorm:"default:false;index" json:"is_deleted"` // soft-deleted by cashier — admin-only, black border, hidden everywhere else
	DeletedReason      string   `json:"deleted_reason"`
	DeletedByName      string   `json:"deleted_by_name"`
	DeletedByRole      string   `json:"deleted_by_role"`
	Archived           bool     `gorm:"default:false;index" json:"archived"` // hidden from order lists but kept for analytics
	ArchiveReason      string   `json:"archive_reason"`
	ReceiptPath        string     `json:"receipt_path"` // uploaded payment receipt photo for online QR payments
	ReceiptPaths       string     `gorm:"type:text" json:"receipt_paths"` // JSON []string — all uploaded receipt photos (split payment support)
	LastReceiptAt      *time.Time `json:"last_receipt_at"`                // last receipt upload time (10-sec rate limit)
	HiddenByUser       bool     `gorm:"default:false;index" json:"hidden_by_user"` // user hid this order from their own history
	BtsTrackingNumber  string   `json:"bts_tracking_number"` // BTS cargo waybill number
	BtsPickupPoint     string   `json:"bts_pickup_point"`    // BTS pickup point name+address (shown to customer)
	BtsBranchLat       float64  `json:"bts_branch_lat"`      // exact BTS branch lat for OSRM routing
	BtsBranchLng       float64  `json:"bts_branch_lng"`      // exact BTS branch lng for OSRM routing
	WorkerNotes        string   `json:"worker_notes"`        // worker comment visible to the customer
	CreatedAt       time.Time   `json:"created_at"`
}

// BtsBranch is a BTS Cargo pickup point configured by admin with verified coordinates.
type BtsBranch struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"not null" json:"name"`     // "БТС — Кургантепа"
	City     string  `gorm:"not null;index" json:"city"` // "Кургантепа"
	Region   string  `json:"region"`                   // "Андижанская область"
	Address  string  `json:"address"`                  // full street address
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	IsActive bool    `gorm:"default:true" json:"is_active"`
}

// Setting is a simple key/value store (e.g. the payment QR image path).
type Setting struct {
	Key   string `gorm:"primaryKey" json:"key"`
	Value string `gorm:"type:text" json:"value"`
}

// WorkerStock is the personal inventory of a worker (pickup point or manager).
// Each worker keeps their own stock per product.
type WorkerStock struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	WorkerID  uint    `gorm:"uniqueIndex:idx_worker_product;not null" json:"worker_id"`
	ProductID uint    `gorm:"uniqueIndex:idx_worker_product;not null" json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `gorm:"default:0" json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ProductComment is a public review/comment left on a specific product.
type ProductComment struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ProductID  uint      `gorm:"index;not null" json:"product_id"`
	AuthorName string    `json:"author_name"`
	Text       string    `gorm:"type:text;not null" json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}

// MarketologPayment records money a marketolog paid back against their debt (transfer/ХР).
type MarketologPayment struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	MarketologID uint      `gorm:"index;not null" json:"marketolog_id"`
	Amount       float64   `gorm:"not null" json:"amount"`
	WorkerID     *uint     `json:"worker_id"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `json:"created_at"`
}

type OrderItem struct {
	ID               uint    `gorm:"primaryKey" json:"id"`
	OrderID          uint    `gorm:"not null" json:"order_id"`
	ProductID        uint    `gorm:"not null" json:"product_id"`
	Product          Product `json:"product"`
	Quantity         int     `gorm:"not null" json:"quantity"`
	OriginalQuantity int     `gorm:"default:0" json:"original_quantity"`
	UnitType         string  `gorm:"default:'pack';not null" json:"unit_type"`
	Price            float64 `gorm:"not null" json:"price"`
}
