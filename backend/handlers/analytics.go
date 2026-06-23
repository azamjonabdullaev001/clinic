package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// parsePaymentSplits decodes the order's stored JSON payment breakdown (empty when the
// sale was paid via a single method).
func parsePaymentSplits(raw string) []PaymentSplit {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	var splits []PaymentSplit
	if err := json.Unmarshal([]byte(raw), &splits); err != nil {
		return nil
	}
	return splits
}

type AnalyticsPoint struct {
	Label   string  `json:"label"`
	Revenue float64 `json:"revenue"`
	Orders  int     `json:"orders"`
}

type TopProduct struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	TotalQty    int     `json:"total_qty"`
	TotalPacks  int     `json:"total_packs"`
	Revenue     float64 `json:"revenue"`
}

type DoctorReferral struct {
	DoctorName  string  `json:"doctor_name"`
	OrderCount  int     `json:"order_count"`
	Capsules    int     `json:"capsules"`
	Pieces      int     `json:"pieces"`
	TotalRevenue float64 `json:"total_revenue"`
}

// CategoryStat is a non-overlapping summary bucket (own patients / doctors / regular / marketolog).
type CategoryStat struct {
	Orders   int     `json:"orders"`
	Capsules int     `json:"capsules"`
	Pieces   int     `json:"pieces"`
	Revenue  float64 `json:"revenue"`
}

type MarketologProduct struct {
	ProductName string  `json:"product_name"`
	Capsules    int     `json:"capsules"`
	Pieces      int     `json:"pieces"`
	Revenue     float64 `json:"revenue"`
}

type MarketologStat struct {
	TotalRevenue  float64             `json:"total_revenue"`
	TotalOrders   int                 `json:"total_orders"`
	TotalCapsules int                 `json:"total_capsules"`
	TotalPieces   int                 `json:"total_pieces"`
	Products      []MarketologProduct `json:"products"`
}

// DiscountSummary aggregates cashier-typed discounts across all offline sales
// in the period, so the admin sees how much was given away.
type DiscountSummary struct {
	OrderCount     int     `json:"order_count"`     // orders that had a discount
	MinPercent     float64 `json:"min_percent"`     // smallest non-zero discount used
	MaxPercent     float64 `json:"max_percent"`     // largest discount used
	AvgPercent     float64 `json:"avg_percent"`     // average across discounted orders
	TotalDiscount  float64 `json:"total_discount"`  // sum of money discounted (estimated from final price)
}

type AnalyticsResponse struct {
	Points          []AnalyticsPoint         `json:"points"`
	TopProducts     []TopProduct             `json:"top_products"`
	DoctorReferrals []DoctorReferral         `json:"doctor_referrals"`
	TotalRevenue    float64                  `json:"total_revenue"`
	TotalOrders     int                      `json:"total_orders"`
	Marketolog      MarketologStat           `json:"marketolog"`
	VIP             MarketologStat           `json:"vip"`
	Breakdown       map[string]*CategoryStat `json:"breakdown"`
	Discounts       DiscountSummary          `json:"discounts"`
}

func GetAnalytics(c *gin.Context) {
	period := c.Query("period") // daily, weekly, monthly, custom
	dateStr := c.Query("date")  // YYYY-MM-DD for custom

	loc := time.Local
	now := time.Now().In(loc)

	var startTime, endTime time.Time
	var points []AnalyticsPoint

	switch period {
	case "weekly":
		// 14 points, each = 12 hours, covering last 7 days
		endTime = now
		startTime = now.AddDate(0, 0, -7)
		for i := 0; i < 14; i++ {
			t := startTime.Add(time.Duration(i) * 12 * time.Hour)
			points = append(points, AnalyticsPoint{
				Label: t.Format("02.01 15:04"),
			})
		}
	case "monthly":
		// 30 points, each = 1 day
		endTime = now
		startTime = now.AddDate(0, 0, -29)
		startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, loc)
		for i := 0; i < 30; i++ {
			t := startTime.AddDate(0, 0, i)
			points = append(points, AnalyticsPoint{
				Label: t.Format("02.01"),
			})
		}
		endTime = startTime.AddDate(0, 0, 30)
	case "yearly":
		// 12 points, one per month, ending with the current month so today is included.
		curMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc)
		startTime = curMonth.AddDate(0, -11, 0)
		endTime = curMonth.AddDate(0, 1, 0)
		for i := 0; i < 12; i++ {
			t := startTime.AddDate(0, i, 0)
			points = append(points, AnalyticsPoint{Label: t.Format("01.2006")})
		}
	case "custom":
		// 24 points for a specific day
		var chosenDay time.Time
		if dateStr != "" {
			parsed, err := time.ParseInLocation("2006-01-02", dateStr, loc)
			if err == nil {
				chosenDay = parsed
			} else {
				chosenDay = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
			}
		} else {
			chosenDay = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		}
		startTime = chosenDay
		endTime = chosenDay.Add(24 * time.Hour)
		for h := 0; h < 24; h++ {
			t := chosenDay.Add(time.Duration(h) * time.Hour)
			points = append(points, AnalyticsPoint{
				Label: t.Format("15:04"),
			})
		}
	default:
		// daily: 24 points, one per hour for today
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		startTime = today
		endTime = today.Add(24 * time.Hour)
		for h := 0; h < 24; h++ {
			t := today.Add(time.Duration(h) * time.Hour)
			points = append(points, AnalyticsPoint{
				Label: t.Format("15:04"),
			})
		}
	}

	// Stream all in-range orders in batches and aggregate in a single pass, so peak
	// memory stays bounded no matter how many orders exist (100k / 500k / 1M+). Every
	// figure below is computed exactly as the previous load-everything version did —
	// only the iteration is batched instead of materialising the whole table at once.
	type ProductStat struct {
		ProductID   uint
		ProductName string
		TotalQty    int
		TotalPacks  int
		Revenue     float64
	}

	var marketolog MarketologStat
	mktMap := make(map[uint]*MarketologProduct)
	statsMap := make(map[uint]*ProductStat)
	// Own patient > doctor referral > regular; marketolog mirrored in after the pass.
	breakdown := map[string]*CategoryStat{
		"vip":     {},
		"doctor":  {},
		"regular": {},
	}
	var vip MarketologStat
	vipMap := make(map[uint]*MarketologProduct)
	doctorMap := make(map[string]*DoctorReferral)
	var disc DiscountSummary
	sumPct := 0.0
	nonMktOrders := 0

	var batch []models.Order
	database.DB.Where("created_at >= ? AND created_at < ? AND status != ? AND is_deleted = ?", startTime, endTime, "cancelled", false).
		Preload("Items.Product").
		FindInBatches(&batch, 1000, func(tx *gorm.DB, _ int) error {
			for bi := range batch {
				order := batch[bi]

				// Cashier-typed discount stats across every order in range (including marketolog).
				if order.DiscountPercent > 0 {
					var orderRevenue float64
					for _, it := range order.Items {
						orderRevenue += it.Price
					}
					// orderRevenue already reflects the discount, so the original was revenue / (1 - pct/100).
					originalTotal := orderRevenue
					if order.DiscountPercent < 100 {
						originalTotal = orderRevenue / (1 - order.DiscountPercent/100)
					}
					disc.OrderCount++
					sumPct += order.DiscountPercent
					if disc.MinPercent == 0 || order.DiscountPercent < disc.MinPercent {
						disc.MinPercent = order.DiscountPercent
					}
					if order.DiscountPercent > disc.MaxPercent {
						disc.MaxPercent = order.DiscountPercent
					}
					disc.TotalDiscount += originalTotal - orderRevenue
				}

				// Marketolog (marketplace) sales are a separate "debt", kept out of the main
				// revenue / top-products / breakdown figures.
				if order.MarketologID != nil {
					marketolog.TotalOrders++
					for _, item := range order.Items {
						if item.Quantity <= 0 {
							continue
						}
						marketolog.TotalRevenue += item.Price
						m, ok := mktMap[item.ProductID]
						if !ok {
							m = &MarketologProduct{ProductName: item.Product.Name}
							mktMap[item.ProductID] = m
						}
						m.Revenue += item.Price
						if item.UnitType == "piece" {
							marketolog.TotalPieces += item.Quantity
							m.Pieces += item.Quantity
						} else {
							marketolog.TotalCapsules += item.Quantity
							m.Capsules += item.Quantity
						}
					}
					continue
				}

				nonMktOrders++

				orderTime := order.CreatedAt.In(loc)
				// revenueAll counts every line (incl. items zeroed-out at the till) — used for
				// the chart points/total, exactly as the original points loop did. revenueActive
				// counts only kept (qty>0) lines — used for the category/doctor breakdown.
				var revenueAll, revenueActive float64
				caps, pcs := 0, 0
				for _, item := range order.Items {
					revenueAll += item.Price
					if item.Quantity <= 0 {
						continue
					}
					revenueActive += item.Price
					if item.UnitType == "piece" {
						pcs += item.Quantity
					} else {
						caps += item.Quantity
					}
				}

				// Chart points: add revenue/order count to the matching time bucket.
				switch period {
				case "weekly":
					for pi := range points {
						slotStart := startTime.Add(time.Duration(pi) * 12 * time.Hour)
						slotEnd := slotStart.Add(12 * time.Hour)
						if !orderTime.Before(slotStart) && orderTime.Before(slotEnd) {
							points[pi].Revenue += revenueAll
							points[pi].Orders++
							break
						}
					}
				case "monthly":
					for pi := range points {
						slotStart := startTime.AddDate(0, 0, pi)
						slotEnd := slotStart.AddDate(0, 0, 1)
						if !orderTime.Before(slotStart) && orderTime.Before(slotEnd) {
							points[pi].Revenue += revenueAll
							points[pi].Orders++
							break
						}
					}
				case "yearly":
					monthIndex := (orderTime.Year()-startTime.Year())*12 + int(orderTime.Month()) - int(startTime.Month())
					if monthIndex >= 0 && monthIndex < len(points) {
						points[monthIndex].Revenue += revenueAll
						points[monthIndex].Orders++
					}
				default:
					// hourly (daily or custom)
					hour := orderTime.Hour()
					if hour >= 0 && hour < len(points) {
						points[hour].Revenue += revenueAll
						points[hour].Orders++
					}
				}

				// Top products (kept lines only).
				for _, item := range order.Items {
					if item.Quantity <= 0 {
						continue // skip items removed at the till
					}
					s, ok := statsMap[item.ProductID]
					if !ok {
						s = &ProductStat{ProductID: item.ProductID, ProductName: item.Product.Name}
						statsMap[item.ProductID] = s
					}
					s.Revenue += item.Price
					if item.UnitType == "piece" {
						s.TotalQty += item.Quantity
					} else {
						s.TotalPacks += item.Quantity
						s.TotalQty += item.Quantity * item.Product.QuantityPerPack
					}
				}

				// Category breakdown + VIP retail value + doctor referrals.
				cat := "regular"
				if order.IsVIP {
					cat = "vip"
				} else if strings.TrimSpace(order.ReferredBy) != "" {
					cat = "doctor"
				}
				b := breakdown[cat]
				b.Orders++
				b.Capsules += caps
				b.Pieces += pcs
				b.Revenue += revenueActive

				switch cat {
				case "vip":
					// Own-patient sales are free (item price 0), so report the retail value
					// of the medicine handed out rather than the recorded 0.
					vip.TotalOrders++
					vip.TotalCapsules += caps
					vip.TotalPieces += pcs
					vipValue := 0.0
					for _, item := range order.Items {
						if item.Quantity <= 0 {
							continue
						}
						item.Product.ComputePackPrice()
						v := item.Product.PricePerPack * float64(item.Quantity)
						if item.UnitType == "piece" {
							v = item.Product.PricePerPill * float64(item.Quantity)
						}
						vipValue += v
						m, ok := vipMap[item.ProductID]
						if !ok {
							m = &MarketologProduct{ProductName: item.Product.Name}
							vipMap[item.ProductID] = m
						}
						m.Revenue += v
						if item.UnitType == "piece" {
							m.Pieces += item.Quantity
						} else {
							m.Capsules += item.Quantity
						}
					}
					vip.TotalRevenue += vipValue
					b.Revenue += vipValue
				case "doctor":
					d, ok := doctorMap[order.ReferredBy]
					if !ok {
						d = &DoctorReferral{DoctorName: order.ReferredBy}
						doctorMap[order.ReferredBy] = d
					}
					d.OrderCount++
					d.Capsules += caps
					d.Pieces += pcs
					d.TotalRevenue += revenueActive
				}
			}
			return nil
		})

	// ---- finalize the streamed aggregates ----
	for _, m := range mktMap {
		marketolog.Products = append(marketolog.Products, *m)
	}
	// Mirror the marketolog debt into the summary so the admin sees every channel separately.
	breakdown["marketolog"] = &CategoryStat{
		Orders:   marketolog.TotalOrders,
		Capsules: marketolog.TotalCapsules,
		Pieces:   marketolog.TotalPieces,
		Revenue:  marketolog.TotalRevenue,
	}

	// Sort products by revenue descending, take top 10.
	var topSlice []TopProduct
	for _, s := range statsMap {
		topSlice = append(topSlice, TopProduct{
			ProductID:   s.ProductID,
			ProductName: s.ProductName,
			TotalQty:    s.TotalQty,
			TotalPacks:  s.TotalPacks,
			Revenue:     s.Revenue,
		})
	}
	for i := 0; i < len(topSlice)-1; i++ {
		for j := i + 1; j < len(topSlice); j++ {
			if topSlice[j].Revenue > topSlice[i].Revenue {
				topSlice[i], topSlice[j] = topSlice[j], topSlice[i]
			}
		}
	}
	if len(topSlice) > 10 {
		topSlice = topSlice[:10]
	}

	var totalRevenue float64
	for _, p := range points {
		totalRevenue += p.Revenue
	}

	for _, m := range vipMap {
		vip.Products = append(vip.Products, *m)
	}
	var doctorSlice []DoctorReferral
	for _, d := range doctorMap {
		doctorSlice = append(doctorSlice, *d)
	}
	for i := 0; i < len(doctorSlice)-1; i++ {
		for j := i + 1; j < len(doctorSlice); j++ {
			if doctorSlice[j].TotalRevenue > doctorSlice[i].TotalRevenue {
				doctorSlice[i], doctorSlice[j] = doctorSlice[j], doctorSlice[i]
			}
		}
	}

	if disc.OrderCount > 0 {
		disc.AvgPercent = sumPct / float64(disc.OrderCount)
	}

	c.JSON(http.StatusOK, AnalyticsResponse{
		Points:          points,
		TopProducts:     topSlice,
		DoctorReferrals: doctorSlice,
		TotalRevenue:    totalRevenue,
		TotalOrders:     nonMktOrders,
		Marketolog:      marketolog,
		VIP:             vip,
		Breakdown:       breakdown,
		Discounts:       disc,
	})
}

// marketologAnalytics computes debt analytics for one marketolog over a period.
func marketologAnalytics(marketologID uint, period, dateStr string) gin.H {
	loc := time.Local
	startTime, endTime, points := buildAnalyticsPeriod(period, dateStr)

	var orders []models.Order
	database.DB.Where("marketolog_id = ? AND status != ? AND is_deleted = ? AND created_at >= ? AND created_at < ?",
		marketologID, "cancelled", false, startTime, endTime).
		Preload("Items.Product").
		Find(&orders)

	var totalRevenue float64
	capsules, pieces := 0, 0
	prodMap := map[uint]*MarketologProduct{}
	for _, order := range orders {
		var revenue float64
		for _, item := range order.Items {
			if item.Quantity <= 0 {
				continue
			}
			revenue += item.Price
			p, ok := prodMap[item.ProductID]
			if !ok {
				p = &MarketologProduct{ProductName: item.Product.Name}
				prodMap[item.ProductID] = p
			}
			p.Revenue += item.Price
			if item.UnitType == "piece" {
				pieces += item.Quantity
				p.Pieces += item.Quantity
			} else {
				capsules += item.Quantity
				p.Capsules += item.Quantity
			}
		}
		totalRevenue += revenue
		addRevenueToPoint(points, period, startTime, order.CreatedAt.In(loc), revenue)
	}
	var prods []MarketologProduct
	for _, p := range prodMap {
		prods = append(prods, *p)
	}

	return gin.H{
		"points":         points,
		"total_revenue":  totalRevenue,
		"total_orders":   len(orders),
		"total_capsules": capsules,
		"total_pieces":   pieces,
		"products":       prods,
	}
}

// GetMarketologOwnAnalytics returns the logged-in marketolog's debt analytics.
func GetMarketologOwnAnalytics(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	c.JSON(http.StatusOK, marketologAnalytics(workerID.(uint), c.Query("period"), c.Query("date")))
}

// GetMarketologStatsAdmin returns one marketolog's debt analytics for the admin.
func GetMarketologStatsAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var worker models.Worker
	if database.DB.First(&worker, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Маркетолог не найден"})
		return
	}
	res := marketologAnalytics(uint(id), c.Query("period"), c.Query("date"))
	res["marketolog"] = gin.H{"id": worker.ID, "name": worker.Name}
	c.JSON(http.StatusOK, res)
}

// buildAnalyticsPeriod returns the time range and empty chart points for a period.
func buildAnalyticsPeriod(period, dateStr string) (time.Time, time.Time, []AnalyticsPoint) {
	loc := time.Local
	now := time.Now().In(loc)
	var startTime, endTime time.Time
	var points []AnalyticsPoint

	switch period {
	case "weekly":
		startTime = now.AddDate(0, 0, -7)
		endTime = now
		for i := 0; i < 14; i++ {
			t := startTime.Add(time.Duration(i) * 12 * time.Hour)
			points = append(points, AnalyticsPoint{Label: t.Format("02.01 15:04")})
		}
	case "monthly":
		startTime = now.AddDate(0, 0, -29)
		startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, loc)
		for i := 0; i < 30; i++ {
			t := startTime.AddDate(0, 0, i)
			points = append(points, AnalyticsPoint{Label: t.Format("02.01")})
		}
		endTime = startTime.AddDate(0, 0, 30)
	case "yearly":
		// 12 points, one per month, ending with the current month so today is included.
		curMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc)
		startTime = curMonth.AddDate(0, -11, 0)
		endTime = curMonth.AddDate(0, 1, 0)
		for i := 0; i < 12; i++ {
			t := startTime.AddDate(0, i, 0)
			points = append(points, AnalyticsPoint{Label: t.Format("01.2006")})
		}
	case "custom":
		chosenDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		if dateStr != "" {
			if parsed, err := time.ParseInLocation("2006-01-02", dateStr, loc); err == nil {
				chosenDay = parsed
			}
		}
		startTime = chosenDay
		endTime = chosenDay.Add(24 * time.Hour)
		for h := 0; h < 24; h++ {
			points = append(points, AnalyticsPoint{Label: chosenDay.Add(time.Duration(h) * time.Hour).Format("15:04")})
		}
	default:
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		startTime = today
		endTime = today.Add(24 * time.Hour)
		for h := 0; h < 24; h++ {
			points = append(points, AnalyticsPoint{Label: today.Add(time.Duration(h) * time.Hour).Format("15:04")})
		}
	}
	return startTime, endTime, points
}

// addRevenueToPoint places an order's revenue into the correct chart slot.
func addRevenueToPoint(points []AnalyticsPoint, period string, startTime, orderTime time.Time, revenue float64) {
	switch period {
	case "weekly":
		for pi := range points {
			slotStart := startTime.Add(time.Duration(pi) * 12 * time.Hour)
			slotEnd := slotStart.Add(12 * time.Hour)
			if !orderTime.Before(slotStart) && orderTime.Before(slotEnd) {
				points[pi].Revenue += revenue
				points[pi].Orders++
				return
			}
		}
	case "monthly":
		for pi := range points {
			slotStart := startTime.AddDate(0, 0, pi)
			slotEnd := slotStart.AddDate(0, 0, 1)
			if !orderTime.Before(slotStart) && orderTime.Before(slotEnd) {
				points[pi].Revenue += revenue
				points[pi].Orders++
				return
			}
		}
	case "yearly":
		monthIndex := (orderTime.Year()-startTime.Year())*12 + int(orderTime.Month()) - int(startTime.Month())
		if monthIndex >= 0 && monthIndex < len(points) {
			points[monthIndex].Revenue += revenue
			points[monthIndex].Orders++
		}
	default:
		hour := orderTime.Hour()
		if hour >= 0 && hour < len(points) {
			points[hour].Revenue += revenue
			points[hour].Orders++
		}
	}
}

// GetWorkerAnalytics returns the logged-in worker's own report (orders they
// delivered/created and revenue) for the requested period.
func GetWorkerAnalytics(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	period := c.Query("period")
	dateStr := c.Query("date")
	loc := time.Local

	startTime, endTime, points := buildAnalyticsPeriod(period, dateStr)

	var orders []models.Order
	database.DB.Where("worker_id = ? AND status = ? AND is_deleted = ? AND created_at >= ? AND created_at < ?",
		workerID, "delivered", false, startTime, endTime).
		Preload("Items.Product").
		Find(&orders)

	// Resolve marketolog id -> name for the breakdown.
	mktNames := map[uint]string{}
	var mgrs []models.Worker
	database.DB.Where("role = ?", "manager").Find(&mgrs)
	for _, w := range mgrs {
		mktNames[w.ID] = w.Name
	}

	var totalRevenue float64
	createdCount := 0
	confirmedCount := 0

	type catAgg struct {
		Orders   int     `json:"orders"`
		Capsules int     `json:"capsules"`
		Pieces   int     `json:"pieces"`
		Revenue  float64 `json:"revenue"`
	}
	cats := map[string]*catAgg{
		"vip":        {},
		"marketolog": {},
		"doctor":     {},
		"regular":    {},
	}
	byDoctor := map[string]*catAgg{}
	byMarketolog := map[string]*catAgg{}
	// Payment-method breakdown. Key is `card_type` when payment is "card", otherwise the method name.
	byPayment := map[string]*catAgg{}

	mainOrders := 0
	for _, order := range orders {
		var revenue, caps, pcs = 0.0, 0, 0
		for _, item := range order.Items {
			if item.Quantity <= 0 {
				continue
			}
			revenue += item.Price
			if item.UnitType == "piece" {
				pcs += item.Quantity
			} else {
				caps += item.Quantity
			}
		}

		// Classify into a single category.
		cat := "regular"
		if order.MarketologID != nil {
			cat = "marketolog"
		} else if order.IsVIP {
			cat = "vip"
		} else if strings.TrimSpace(order.ReferredBy) != "" {
			cat = "doctor"
		}

		// Marketolog (debt) sales never mix into the main revenue/points/totals.
		if cat != "marketolog" {
			totalRevenue += revenue
			mainOrders++
			addRevenueToPoint(points, period, startTime, order.CreatedAt.In(loc), revenue)
			if order.IsOffline && !order.IsNurseOrder {
				createdCount++
			} else {
				confirmedCount++
			}
			// Bucket by payment method. Split payments distribute their amount across the
			// methods used; single-method sales keep the legacy card-subtype breakout.
			if splits := parsePaymentSplits(order.PaymentSplits); len(splits) > 0 {
				for _, s := range splits {
					if s.Amount == 0 || s.Method == "" {
						continue
					}
					p, ok := byPayment[s.Method]
					if !ok {
						p = &catAgg{}
						byPayment[s.Method] = p
					}
					p.Orders++
					p.Revenue += s.Amount
				}
			} else {
				payKey := order.PaymentMethod
				if payKey == "card" && order.CardType != "" {
					payKey = order.CardType
				}
				if payKey == "" {
					payKey = "free"
				}
				p, ok := byPayment[payKey]
				if !ok {
					p = &catAgg{}
					byPayment[payKey] = p
				}
				p.Orders++
				p.Capsules += caps
				p.Pieces += pcs
				p.Revenue += revenue
			}
		}
		c := cats[cat]
		c.Orders++
		c.Capsules += caps
		c.Pieces += pcs
		c.Revenue += revenue

		if cat == "doctor" {
			d, ok := byDoctor[order.ReferredBy]
			if !ok {
				d = &catAgg{}
				byDoctor[order.ReferredBy] = d
			}
			d.Orders++
			d.Capsules += caps
			d.Pieces += pcs
			d.Revenue += revenue
		}
		if cat == "marketolog" {
			name := mktNames[*order.MarketologID]
			if name == "" {
				name = "Маркетолог"
			}
			m, ok := byMarketolog[name]
			if !ok {
				m = &catAgg{}
				byMarketolog[name] = m
			}
			m.Orders++
			m.Capsules += caps
			m.Pieces += pcs
			m.Revenue += revenue
		}
	}

	type namedAgg struct {
		Name     string  `json:"name"`
		Orders   int     `json:"orders"`
		Capsules int     `json:"capsules"`
		Pieces   int     `json:"pieces"`
		Revenue  float64 `json:"revenue"`
	}
	doctorList := []namedAgg{}
	for name, a := range byDoctor {
		doctorList = append(doctorList, namedAgg{name, a.Orders, a.Capsules, a.Pieces, a.Revenue})
	}
	marketologList := []namedAgg{}
	for name, a := range byMarketolog {
		marketologList = append(marketologList, namedAgg{name, a.Orders, a.Capsules, a.Pieces, a.Revenue})
	}

	// Build top-products list from delivered orders.
	type prodStat struct {
		ProductName string
		Orders      int
		Capsules    int
		Pieces      int
		Revenue     float64
	}
	prodMap := map[uint]*prodStat{}
	for _, order := range orders {
		for _, item := range order.Items {
			if item.Quantity <= 0 {
				continue
			}
			ps, ok := prodMap[item.ProductID]
			if !ok {
				ps = &prodStat{ProductName: item.Product.Name}
				prodMap[item.ProductID] = ps
			}
			ps.Orders++
			ps.Revenue += item.Price
			if item.UnitType == "piece" {
				ps.Pieces += item.Quantity
			} else {
				ps.Capsules += item.Quantity
			}
		}
	}
	type prodStatOut struct {
		ProductID   uint    `json:"product_id"`
		ProductName string  `json:"product_name"`
		Orders      int     `json:"orders"`
		Capsules    int     `json:"capsules"`
		Pieces      int     `json:"pieces"`
		Revenue     float64 `json:"revenue"`
	}
	var topProducts []prodStatOut
	for id, ps := range prodMap {
		topProducts = append(topProducts, prodStatOut{id, ps.ProductName, ps.Orders, ps.Capsules, ps.Pieces, ps.Revenue})
	}
	for i := 0; i < len(topProducts)-1; i++ {
		for j := i + 1; j < len(topProducts); j++ {
			if topProducts[j].Revenue > topProducts[i].Revenue {
				topProducts[i], topProducts[j] = topProducts[j], topProducts[i]
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"points":          points,
		"total_revenue":   totalRevenue,
		"total_orders":    mainOrders,
		"created_count":   createdCount,
		"confirmed_count": confirmedCount,
		"breakdown":       cats,
		"by_doctor":       doctorList,
		"by_marketolog":   marketologList,
		"by_payment":      byPayment,
		"top_products":    topProducts,
	})
}

// ProductChannelStat holds the per-channel totals for a single product.
type ProductChannelStat struct {
	Orders   int `json:"orders"`
	Capsules int `json:"capsules"`
	Pieces   int `json:"pieces"`
}

// GetProductAnalytics returns a breakdown of one product's sales by channel
// (VIP own-patients, online, offline pickup, marketolog).
func GetProductAnalytics(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Препарат не найден"})
		return
	}
	product.ComputePackPrice()

	type ItemRow struct {
		UnitType   string `gorm:"column:unit_type"`
		Quantity   int    `gorm:"column:quantity"`
		IsVIP      bool   `gorm:"column:is_v_ip"`
		IsOffline  bool   `gorm:"column:is_offline"`
		HasMktolog bool   `gorm:"column:has_mktolog"`
	}
	var rows []ItemRow
	database.DB.Table("order_items").
		Select("order_items.unit_type, order_items.quantity, orders.is_v_ip, orders.is_offline, (orders.marketolog_id IS NOT NULL) AS has_mktolog").
		Joins("JOIN orders ON orders.id = order_items.order_id").
		Where("order_items.product_id = ? AND orders.status != ? AND orders.archived = ? AND orders.is_deleted = ?", id, "cancelled", false, false).
		Scan(&rows)

	channels := map[string]*ProductChannelStat{
		"vip":        {},
		"online":     {},
		"offline":    {},
		"marketolog": {},
	}
	totalCaps, totalPcs := 0, 0
	for _, row := range rows {
		var ch string
		if row.IsVIP {
			ch = "vip"
		} else if row.HasMktolog {
			ch = "marketolog"
		} else if row.IsOffline {
			ch = "offline"
		} else {
			ch = "online"
		}
		stat := channels[ch]
		stat.Orders++
		if row.UnitType == "pack" {
			stat.Capsules += row.Quantity
			totalCaps += row.Quantity
		} else {
			stat.Pieces += row.Quantity
			totalPcs += row.Quantity
		}
	}

	totalOrders := 0
	for _, s := range channels {
		totalOrders += s.Orders
	}

	c.JSON(http.StatusOK, gin.H{
		"product_id":     product.ID,
		"product_name":   product.Name,
		"current_stock":  product.StockQuantity,
		"qty_per_pack":   product.QuantityPerPack,
		"total_capsules": totalCaps,
		"total_pieces":   totalPcs,
		"total_orders":   totalOrders,
		"channels":       channels,
	})
}
