package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetDoctors(c *gin.Context) {
	var doctors []models.Doctor
	database.DB.Order("name asc").Find(&doctors)
	c.JSON(http.StatusOK, doctors)
}

func CreateDoctor(c *gin.Context) {
	var input struct {
		Name      string `json:"name" binding:"required"`
		Specialty string `json:"specialty"`
		Phone     string `json:"phone"`
		Password  string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя доктора обязательно"})
		return
	}

	doctor := models.Doctor{Name: input.Name, Specialty: input.Specialty}

	if input.Phone != "" {
		if !phoneRegex.MatchString(input.Phone) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат номера телефона (998XXXXXXXXX)"})
			return
		}
		doctor.Phone = input.Phone
	}

	if input.Password != "" {
		if len(input.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Пароль должен содержать минимум 6 символов"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			return
		}
		doctor.Password = string(hash)
	}

	if err := database.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении доктора"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": doctor.ID, "name": doctor.Name, "specialty": doctor.Specialty,
		"phone": doctor.Phone, "created_at": doctor.CreatedAt,
	})
}

func UpdateDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var doctor models.Doctor
	if err := database.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}
	var input struct {
		Name      string `json:"name"`
		Specialty string `json:"specialty"`
		Phone     string `json:"phone"`
		Password  string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	if input.Name != "" {
		doctor.Name = input.Name
	}
	if input.Specialty != "" {
		doctor.Specialty = input.Specialty
	}
	if input.Phone != "" {
		if !phoneRegex.MatchString(input.Phone) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат номера телефона (998XXXXXXXXX)"})
			return
		}
		doctor.Phone = input.Phone
	}
	if len(input.Password) >= 6 {
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			return
		}
		doctor.Password = string(hash)
	}
	if err := database.DB.Save(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": doctor.ID, "name": doctor.Name, "specialty": doctor.Specialty,
		"phone": doctor.Phone, "created_at": doctor.CreatedAt,
	})
}

func DeleteDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	database.DB.Delete(&models.Doctor{}, id)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

type DoctorStatProduct struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	TotalPacks  int     `json:"total_packs"`
	TotalPieces int     `json:"total_pieces"`
	OrderCount  int     `json:"order_count"`
	Revenue     float64 `json:"revenue"`
}

func GetDoctorStats(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var doctor models.Doctor
	if err := database.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}

	var orders []models.Order
	database.DB.Where("referred_by LIKE ? OR doctor_id = ?", "%"+doctor.Name+"%", doctor.ID).
		Preload("Items.Product").
		Find(&orders)

	statsMap := make(map[uint]*DoctorStatProduct)
	for _, order := range orders {
		for _, item := range order.Items {
			s, ok := statsMap[item.ProductID]
			if !ok {
				s = &DoctorStatProduct{
					ProductID:   item.ProductID,
					ProductName: item.Product.Name,
				}
				statsMap[item.ProductID] = s
			}
			s.Revenue += item.Price
			s.OrderCount++
			if item.UnitType == "piece" {
				s.TotalPieces += item.Quantity
			} else {
				s.TotalPacks += item.Quantity
			}
		}
	}

	var result []DoctorStatProduct
	for _, s := range statsMap {
		result = append(result, *s)
	}
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[j].Revenue > result[i].Revenue {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"doctor":       doctor,
		"total_orders": len(orders),
		"products":     result,
	})
}

// ==================== Doctor Orders ====================

type CreateDoctorOrderInput struct {
	PatientFirstName string             `json:"patient_first_name" binding:"required"`
	PatientLastName  string             `json:"patient_last_name" binding:"required"`
	Items            []OfflineItemInput `json:"items" binding:"required,min=1"`
}

func CreateDoctorOrder(c *gin.Context) {
	doctorID, _ := c.Get("doctorID")
	did := doctorID.(uint)

	var input CreateDoctorOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя, фамилия пациента и препараты обязательны"})
		return
	}

	tx := database.DB.Begin()

	order := models.Order{
		UserID:       nil,
		DoctorID:     &did,
		Status:       "pending",
		Phone:        "offline",
		OrderCode:    generateNurseCode(),
		IsOffline:    true,
		IsNurseOrder: true,
		PatientFName: input.PatientFirstName,
		PatientLName: input.PatientLastName,
		OfflineNote:  input.PatientFirstName + " " + input.PatientLastName,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
		return
	}

	for _, item := range input.Items {
		var product models.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Товар не найден"})
			return
		}
		product.ComputePackPrice()

		unitType := item.UnitType
		if unitType == "" {
			unitType = "pack"
		}

		var price float64
		if unitType == "piece" {
			price = product.PricePerPill * float64(item.Quantity)
		} else {
			price = product.PricePerPack * float64(item.Quantity)
		}

		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitType:  unitType,
			Price:     price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
			return
		}
	}

	tx.Commit()

	database.DB.Preload("Items.Product").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusCreated, order)
}

func GetDoctorOrders(c *gin.Context) {
	doctorID, _ := c.Get("doctorID")
	var orders []models.Order
	database.DB.Where("doctor_id = ? AND is_nurse_order = true", doctorID).
		Preload("Items.Product").
		Order("created_at desc").
		Find(&orders)

	for i := range orders {
		for j := range orders[i].Items {
			orders[i].Items[j].Product.ComputePackPrice()
		}
	}

	c.JSON(http.StatusOK, orders)
}

func GetDoctorProfile(c *gin.Context) {
	doctorID, _ := c.Get("doctorID")
	var doctor models.Doctor
	if err := database.DB.First(&doctor, doctorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": doctor.ID, "name": doctor.Name, "specialty": doctor.Specialty, "phone": doctor.Phone,
	})
}

func GetDoctorAnalytics(c *gin.Context) {
	doctorID, _ := c.Get("doctorID")

	var orders []models.Order
	database.DB.Where("doctor_id = ? AND is_nurse_order = true", doctorID).
		Preload("Items.Product").
		Order("created_at desc").
		Find(&orders)

	statsMap := make(map[uint]*DoctorStatProduct)
	var totalRevenue float64

	for _, order := range orders {
		for _, item := range order.Items {
			item.Product.ComputePackPrice()
			totalRevenue += item.Price
			s, ok := statsMap[item.ProductID]
			if !ok {
				s = &DoctorStatProduct{
					ProductID:   item.ProductID,
					ProductName: item.Product.Name,
				}
				statsMap[item.ProductID] = s
			}
			s.Revenue += item.Price
			s.OrderCount++
			if item.UnitType == "piece" {
				s.TotalPieces += item.Quantity
			} else {
				s.TotalPacks += item.Quantity
			}
		}
	}

	var topProducts []DoctorStatProduct
	for _, s := range statsMap {
		topProducts = append(topProducts, *s)
	}
	for i := 0; i < len(topProducts)-1; i++ {
		for j := i + 1; j < len(topProducts); j++ {
			if topProducts[j].Revenue > topProducts[i].Revenue {
				topProducts[i], topProducts[j] = topProducts[j], topProducts[i]
			}
		}
	}
	if len(topProducts) > 10 {
		topProducts = topProducts[:10]
	}

	recentOrders := orders
	if len(recentOrders) > 20 {
		recentOrders = recentOrders[:20]
	}

	c.JSON(http.StatusOK, gin.H{
		"total_orders":  len(orders),
		"total_revenue": totalRevenue,
		"top_products":  topProducts,
		"recent_orders": recentOrders,
	})
}
