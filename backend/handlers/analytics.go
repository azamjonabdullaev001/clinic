package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
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
	TotalRevenue float64 `json:"total_revenue"`
}

type AnalyticsResponse struct {
	Points          []AnalyticsPoint `json:"points"`
	TopProducts     []TopProduct     `json:"top_products"`
	DoctorReferrals []DoctorReferral `json:"doctor_referrals"`
	TotalRevenue    float64          `json:"total_revenue"`
	TotalOrders     int              `json:"total_orders"`
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
	var orders []models.Order
	database.DB.Where("created_at >= ? AND created_at < ? AND status != ?", startTime, endTime, "cancelled").
		Preload("Items.Product").
		Find(&orders)

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

	// Doctor referral stats
	doctorMap := make(map[string]*DoctorReferral)
	for _, order := range orders {
		if order.ReferredBy == "" {
			continue
		}
		d, ok := doctorMap[order.ReferredBy]
		if !ok {
			d = &DoctorReferral{DoctorName: order.ReferredBy}
			doctorMap[order.ReferredBy] = d
		}
		d.OrderCount++
		for _, item := range order.Items {
			d.TotalRevenue += item.Price
		}
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
	})
}
