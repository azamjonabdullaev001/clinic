package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

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

type AnalyticsResponse struct {
	Points          []AnalyticsPoint         `json:"points"`
	TopProducts     []TopProduct             `json:"top_products"`
	DoctorReferrals []DoctorReferral         `json:"doctor_referrals"`
	TotalRevenue    float64                  `json:"total_revenue"`
	TotalOrders     int                      `json:"total_orders"`
	Marketolog      MarketologStat           `json:"marketolog"`
	VIP             MarketologStat           `json:"vip"`
	Breakdown       map[string]*CategoryStat `json:"breakdown"`
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

	// Fetch all delivered/completed orders in range
	var allOrders []models.Order
	database.DB.Where("created_at >= ? AND created_at < ? AND status != ?", startTime, endTime, "cancelled").
		Preload("Items.Product").
		Find(&allOrders)

	// Marketolog (marketplace) sales are treated as a separate "debt" and are kept
	// out of the main revenue/top-products figures.
	var orders []models.Order
	var mktOrders []models.Order
	for _, o := range allOrders {
		if o.MarketologID != nil {
			mktOrders = append(mktOrders, o)
		} else {
			orders = append(orders, o)
		}
	}

	// Aggregate the marketolog debt section.
	var marketolog MarketologStat
	marketolog.TotalOrders = len(mktOrders)
	mktMap := make(map[uint]*MarketologProduct)
	for _, order := range mktOrders {
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
	}
	for _, m := range mktMap {
		marketolog.Products = append(marketolog.Products, *m)
	}

	// Fill points with revenue/order counts
	for _, order := range orders {
		orderTime := order.CreatedAt.In(loc)
		var revenue float64
		for _, item := range order.Items {
			revenue += item.Price
		}

		switch period {
		case "weekly":
			for pi := range points {
				slotStart := startTime.Add(time.Duration(pi) * 12 * time.Hour)
				slotEnd := slotStart.Add(12 * time.Hour)
				if !orderTime.Before(slotStart) && orderTime.Before(slotEnd) {
					points[pi].Revenue += revenue
					points[pi].Orders++
					break
				}
			}
		case "monthly":
			for pi := range points {
				slotStart := startTime.AddDate(0, 0, pi)
				slotEnd := slotStart.AddDate(0, 0, 1)
				if !orderTime.Before(slotStart) && orderTime.Before(slotEnd) {
					points[pi].Revenue += revenue
					points[pi].Orders++
					break
				}
			}
		default:
			// hourly (daily or custom)
			hour := orderTime.Hour()
			if hour >= 0 && hour < len(points) {
				points[hour].Revenue += revenue
				points[hour].Orders++
			}
		}
	}

	// Top 10 products
	type ProductStat struct {
		ProductID   uint
		ProductName string
		TotalQty    int
		TotalPacks  int
		Revenue     float64
	}
	statsMap := make(map[uint]*ProductStat)
	for _, order := range orders {
		for _, item := range order.Items {
			if item.Quantity <= 0 {
				continue // skip items removed at the till
			}
			s, ok := statsMap[item.ProductID]
			if !ok {
				s = &ProductStat{
					ProductID:   item.ProductID,
					ProductName: item.Product.Name,
				}
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
	}

	// Sort by revenue descending, take top 10
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
	// Simple sort (bubble sort for up to 10 products, enough here)
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

	// Classify every non-marketolog order into exactly one bucket (priority:
	// own patient > doctor referral > regular) and aggregate capsule/piece counts.
	// Own-patient (свои пациенты) sales get a product-level breakdown, doctor
	// referrals get per-doctor tablet counts, and the marketolog debt is mirrored
	// into the summary so the admin sees every channel separately.
	breakdown := map[string]*CategoryStat{
		"vip":     {},
		"doctor":  {},
		"regular": {},
		"marketolog": {
			Orders:   marketolog.TotalOrders,
			Capsules: marketolog.TotalCapsules,
			Pieces:   marketolog.TotalPieces,
			Revenue:  marketolog.TotalRevenue,
		},
	}
	var vip MarketologStat
	vipMap := make(map[uint]*MarketologProduct)
	doctorMap := make(map[string]*DoctorReferral)
	for _, order := range orders {
		var revenue float64
		caps, pcs := 0, 0
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
		b.Revenue += revenue

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
			d.TotalRevenue += revenue
		}
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

	c.JSON(http.StatusOK, AnalyticsResponse{
		Points:          points,
		TopProducts:     topSlice,
		DoctorReferrals: doctorSlice,
		TotalRevenue:    totalRevenue,
		TotalOrders:     len(orders),
		Marketolog:      marketolog,
		VIP:             vip,
		Breakdown:       breakdown,
	})
}

// marketologAnalytics computes debt analytics for one marketolog over a period.
func marketologAnalytics(marketologID uint, period, dateStr string) gin.H {
	loc := time.Local
	startTime, endTime, points := buildAnalyticsPeriod(period, dateStr)

	var orders []models.Order
	database.DB.Where("marketolog_id = ? AND status != ? AND created_at >= ? AND created_at < ?",
		marketologID, "cancelled", startTime, endTime).
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
	database.DB.Where("worker_id = ? AND status = ? AND created_at >= ? AND created_at < ?",
		workerID, "delivered", startTime, endTime).
		Preload("Items").
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

	c.JSON(http.StatusOK, gin.H{
		"points":          points,
		"total_revenue":   totalRevenue,
		"total_orders":    mainOrders,
		"created_count":   createdCount,
		"confirmed_count": confirmedCount,
		"breakdown":       cats,
		"by_doctor":       doctorList,
		"by_marketolog":   marketologList,
	})
}
