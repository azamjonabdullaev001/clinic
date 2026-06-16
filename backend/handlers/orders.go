package handlers

import (
	"bytes"
	"clinic-backend/config"
	"clinic-backend/database"
	"clinic-backend/models"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderItemInput struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	UnitType  string `json:"unit_type"`
}

type CreateOrderInput struct {
	Items           []OrderItemInput `json:"items" binding:"required,min=1"`
	Phone           string           `json:"phone" binding:"required"`
	DeliveryAddress string           `json:"delivery_address"`
	Latitude        float64          `json:"latitude"`
	Longitude       float64          `json:"longitude"`
	ReferredBy      string           `json:"referred_by"`
}

func generateOrderCode() string {
	for {
		code := fmt.Sprintf("%06d", rand.Intn(900000)+100000)
		var count int64
		database.DB.Model(&models.Order{}).Where("order_code = ?", code).Count(&count)
		if count == 0 {
			return code
		}
	}
}

func CreateOrder(c *gin.Context) {
	userID, _ := c.Get("userID")

	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные заказа"})
		return
	}

	deliveryAddress := input.DeliveryAddress
	// Require location: either an address text or coordinates
	if deliveryAddress == "" && (input.Latitude == 0 || input.Longitude == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Укажите адрес доставки или выберите точку на карте"})
		return
	}
	// If coordinates provided but no text address, generate a label
	if deliveryAddress == "" && input.Latitude != 0 && input.Longitude != 0 {
		deliveryAddress = fmt.Sprintf("%.6f, %.6f", input.Latitude, input.Longitude)
	}

	uid := userID.(uint)
	order := models.Order{
		UserID: &uid,
		// Online orders start as awaiting_payment: the customer pays via the QR and uploads
		// a receipt, which moves the order to "pending" and reveals it at the pickup point.
		Status:          "awaiting_payment",
		Phone:           input.Phone,
		OrderCode:       generateOrderCode(),
		DeliveryAddress: deliveryAddress,
		Latitude:        input.Latitude,
		Longitude:       input.Longitude,
		ReferredBy:      input.ReferredBy,
	}

	tx := database.DB.Begin()

	// Block the order if the warehouse doesn't have enough of any item.
	if err := checkProductStock(tx, input.Items); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Препарат не найден"})
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
			OrderID:          order.ID,
			ProductID:        item.ProductID,
			Quantity:         item.Quantity,
			OriginalQuantity: item.Quantity,
			UnitType:         unitType,
			Price:            price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
			return
		}
	}

	tx.Commit()
	BroadcastOrders()

	database.DB.Preload("Items.Product").Preload("User").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusCreated, order)
}

func GetUserOrders(c *gin.Context) {
	userID, _ := c.Get("userID")
	var orders []models.Order
	database.DB.Where("user_id = ? AND archived = ? AND is_deleted = ? AND hidden_by_user = ?", userID, false, false, false).
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

// UpdateOrderLocation lets the customer correct the delivery address/coordinates
// for orders that are still pending (not yet in transit or delivered).
func UpdateOrderLocation(c *gin.Context) {
	userID, _ := c.Get("userID")
	orderID := c.Param("id")

	var order models.Order
	if err := database.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}
	if order.UserID == nil || *order.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещён"})
		return
	}
	if order.Status != "pending" && order.Status != "awaiting_payment" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Нельзя изменить адрес: заказ уже передан в доставку"})
		return
	}

	var input struct {
		DeliveryAddress string  `json:"delivery_address"`
		Latitude        float64 `json:"latitude"`
		Longitude       float64 `json:"longitude"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	if input.Latitude == 0 && input.Longitude == 0 && input.DeliveryAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Укажите адрес или координаты"})
		return
	}

	if err := database.DB.Model(&order).Updates(map[string]interface{}{
		"delivery_address": input.DeliveryAddress,
		"latitude":         input.Latitude,
		"longitude":        input.Longitude,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении"})
		return
	}

	BroadcastOrders()
	c.JSON(http.StatusOK, gin.H{"message": "Адрес доставки обновлён"})
}

func HideUserOrder(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}
	if order.UserID == nil || *order.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Нет доступа"})
		return
	}

	database.DB.Model(&order).Update("hidden_by_user", true)
	c.JSON(http.StatusOK, gin.H{"message": "Заказ скрыт из истории"})
}

func GetOrders(c *gin.Context) {
	limit := 100
	offset := 0
	if l, err := strconv.Atoi(c.DefaultQuery("limit", "100")); err == nil && l > 0 && l <= 500 {
		limit = l
	}
	if o, err := strconv.Atoi(c.DefaultQuery("offset", "0")); err == nil && o >= 0 {
		offset = o
	}

	db := database.DB.Where("archived = ?", false)
	if s := c.Query("status"); s != "" {
		switch s {
		case "edited":
			db = db.Where("is_edited = ? OR is_returned = ?", true, true)
		case "deleted":
			db = db.Where("is_deleted = ?", true)
		default:
			db = db.Where("status = ?", s)
		}
	}
	if df := c.Query("date_from"); df != "" {
		db = db.Where("created_at >= ?", df)
	}
	if dt := c.Query("date_to"); dt != "" {
		db = db.Where("created_at < ?", dt)
	}

	var total int64
	db.Model(&models.Order{}).Count(&total)

	var orders []models.Order
	db.Preload("Items.Product").Preload("User").
		Order("created_at desc").
		Limit(limit).Offset(offset).
		Find(&orders)

	for i := range orders {
		for j := range orders[i].Items {
			orders[i].Items[j].Product.ComputePackPrice()
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный статус"})
		return
	}

	validStatuses := map[string]bool{
		"pending":    true,
		"in_transit": true,
		"delivered":  true,
		"cancelled":  true,
	}
	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный статус"})
		return
	}

	order.Status = input.Status
	database.DB.Save(&order)
	BroadcastOrders()

	database.DB.Preload("Items.Product").Preload("User").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	tx := database.DB.Begin()
	if err := tx.Where("order_id = ?", order.ID).Delete(&models.OrderItem{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении позиций заказа"})
		return
	}
	if err := tx.Delete(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении заказа"})
		return
	}
	tx.Commit()
	BroadcastOrders()

	c.JSON(http.StatusOK, gin.H{"message": "Заказ удалён"})
}

// DeleteAllOrders hides every order from the order lists (admin only). The rows
// stay in the database so analytics still reflects past sales — we only mark them
// archived, optionally recording why.
func DeleteAllOrders(c *gin.Context) {
	var input struct {
		Reason string `json:"reason"`
	}
	c.ShouldBindJSON(&input)

	if err := database.DB.Model(&models.Order{}).
		Where("archived = ?", false).
		Updates(map[string]interface{}{
			"archived":       true,
			"archive_reason": input.Reason,
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении заказов"})
		return
	}
	BroadcastOrders()
	c.JSON(http.StatusOK, gin.H{"message": "Все заказы удалены"})
}

func formatNumber(n float64) string {
	s := fmt.Sprintf("%.0f", n)
	if len(s) <= 3 {
		return s
	}
	var result strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(c)
	}
	return result.String()
}

func sendTelegramNotification(order models.Order) {
	cfg := config.Load()
	if cfg.TelegramBotToken == "" || cfg.TelegramChatID == "" {
		return
	}

	var sb strings.Builder
	sb.WriteString("✅ ТОВАР ВЫДАН\n\n")
	if order.OrderCode != "" {
		sb.WriteString(fmt.Sprintf("🔑 Код заказа: %s\n\n", order.OrderCode))
	}

	var totalSum float64
	for _, item := range order.Items {
		item.Product.ComputePackPrice()
		unitLabel := "коробка"
		if item.UnitType == "piece" {
			unitLabel = "шт"
		}
		sb.WriteString(fmt.Sprintf("💊 Лекарство: %s\n", item.Product.Name))
		sb.WriteString(fmt.Sprintf("📦 Количество: %d %s\n", item.Quantity, unitLabel))
		sb.WriteString(fmt.Sprintf("💰 Сумма: %s сум\n\n", formatNumber(item.Price)))
		totalSum += item.Price
	}

	if len(order.Items) > 1 {
		sb.WriteString(fmt.Sprintf("💰 Общая сумма: %s сум\n\n", formatNumber(totalSum)))
	}

	sb.WriteString(fmt.Sprintf("👤 Имя: %s\n", order.User.FirstName))
	sb.WriteString(fmt.Sprintf("👨 Фамилия: %s\n", order.User.LastName))
	if order.User.MiddleName != "" {
		sb.WriteString(fmt.Sprintf("👤 Отчество: %s\n", order.User.MiddleName))
	}
	sb.WriteString(fmt.Sprintf("📱 Телефон: +%s\n", order.Phone))
	sb.WriteString("🌍 Страна: O'zbekiston")

	payload := map[string]string{
		"chat_id": cfg.TelegramChatID,
		"text":    sb.String(),
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Telegram marshal error: %v", err)
		return
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.TelegramBotToken)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Telegram send error: %v", err)
		return
	}
	defer resp.Body.Close()
}

// UploadOrderReceipt handles payment receipt upload for online orders
// CRITICAL: Only moves order from "awaiting_payment" to "pending" after receipt upload
// This makes the order VISIBLE to pickup points ONLY after payment proof
func UploadOrderReceipt(c *gin.Context) {
	userID, _ := c.Get("userID")
	orderID := c.Param("id")

	// Get order and verify ownership
	var order models.Order
	if err := database.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	if order.UserID == nil || *order.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещён"})
		return
	}

	// Only awaiting_payment orders can receive receipts
	if order.Status != "awaiting_payment" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Этот заказ не ожидает оплаты"})
		return
	}

	// Parse primary receipt file
	file, err := c.FormFile("receipt")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл чека не найден"})
		return
	}
	if file.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл слишком большой (максимум 10MB)"})
		return
	}
	allowedReceiptTypes := map[string]string{
		"image/jpeg": ".jpg", "image/jpg": ".jpg",
		"image/png": ".png", "image/gif": ".gif", "image/webp": ".webp",
	}
	if _, ok := allowedReceiptTypes[file.Header.Get("Content-Type")]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл должен быть изображением (jpg/png/gif/webp)"})
		return
	}

	uploadsDir := "uploads/receipts"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}

	makeReceiptPath := func(f *multipart.FileHeader, idx int) (string, bool) {
		ts := fmt.Sprintf("%d", time.Now().UnixNano())
		ct := f.Header.Get("Content-Type")
		ext := allowedReceiptTypes[ct]
		if ext == "" {
			ext = strings.ToLower(filepath.Ext(f.Filename))
			if ext == "" || len(ext) > 6 {
				ext = ".jpg"
			}
		}
		fname := fmt.Sprintf("%d-%s-%d%s", order.ID, ts, idx, ext)
		fpath := fmt.Sprintf("%s/%s", uploadsDir, fname)
		if err := c.SaveUploadedFile(f, fpath); err != nil {
			return "", false
		}
		return fpath, true
	}

	firstPath, ok := makeReceiptPath(file, 0)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}

	allPaths := []string{firstPath}

	// Also accept additional files sent as `receipts[]`
	form, _ := c.MultipartForm()
	if form != nil {
		for i, extra := range form.File["receipts"] {
			if extra.Size > 10*1024*1024 {
				continue
			}
			if !strings.Contains(extra.Header.Get("Content-Type"), "image/") {
				continue
			}
			if p, ok := makeReceiptPath(extra, i+1); ok {
				allPaths = append(allPaths, p)
			}
		}
	}

	pathsJSON, _ := json.Marshal(allPaths)
	now := time.Now()

	if err := database.DB.Model(&order).Updates(map[string]interface{}{
		"receipt_path":    firstPath,
		"receipt_paths":   string(pathsJSON),
		"last_receipt_at": now,
		"status":          "pending",
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении заказа"})
		return
	}

	BroadcastOrders()

	c.JSON(http.StatusOK, gin.H{
		"message":        "Чек успешно загружен",
		"receipt_path":   firstPath,
		"receipt_paths":  allPaths,
		"receipts_count": len(allPaths),
	})
}

// AddOrderReceipt adds another payment receipt to an already-pending online order.
// Supports split payment: customer pays from multiple bank accounts and uploads each receipt.
// Rate-limited to one upload per 10 seconds to prevent abuse.
func AddOrderReceipt(c *gin.Context) {
	userID, _ := c.Get("userID")
	orderID := c.Param("id")

	var order models.Order
	if err := database.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	if order.UserID == nil || *order.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещён"})
		return
	}

	if order.Status != "pending" && order.Status != "awaiting_payment" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Нельзя добавить чек к этому заказу"})
		return
	}

	// Rate limit: 10 seconds between receipt uploads
	if order.LastReceiptAt != nil {
		elapsed := time.Since(*order.LastReceiptAt)
		if elapsed < 10*time.Second {
			remaining := int(10-elapsed.Seconds()) + 1
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       fmt.Sprintf("Подождите %d сек. перед отправкой следующего чека", remaining),
				"retry_after": remaining,
			})
			return
		}
	}

	file, err := c.FormFile("receipt")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл чека не найден"})
		return
	}
	if file.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл слишком большой (максимум 10MB)"})
		return
	}
	addReceiptTypes := map[string]string{
		"image/jpeg": ".jpg", "image/jpg": ".jpg",
		"image/png": ".png", "image/gif": ".gif", "image/webp": ".webp",
	}
	addExt, typeOk := addReceiptTypes[file.Header.Get("Content-Type")]
	if !typeOk {
		addExt = strings.ToLower(filepath.Ext(file.Filename))
		if addExt == "" || len(addExt) > 6 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Файл должен быть изображением (jpg/png/gif/webp)"})
			return
		}
	}

	uploadsDir := "uploads/receipts"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}

	filename := fmt.Sprintf("%d-%d%s", order.ID, time.Now().UnixNano(), addExt)
	fpath := fmt.Sprintf("%s/%s", uploadsDir, filename)

	if err := c.SaveUploadedFile(file, fpath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}

	// Parse existing receipt_paths or fall back to receipt_path
	var paths []string
	if order.ReceiptPaths != "" {
		_ = json.Unmarshal([]byte(order.ReceiptPaths), &paths)
	}
	if len(paths) == 0 && order.ReceiptPath != "" {
		paths = []string{order.ReceiptPath}
	}
	paths = append(paths, fpath)
	pathsJSON, _ := json.Marshal(paths)

	now := time.Now()
	updates := map[string]interface{}{
		"receipt_paths":   string(pathsJSON),
		"last_receipt_at": now,
	}
	// If first receipt was missed somehow and order still awaiting — fix it
	if order.Status == "awaiting_payment" {
		updates["status"] = "pending"
		if order.ReceiptPath == "" {
			updates["receipt_path"] = fpath
		}
	}

	if err := database.DB.Model(&order).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении заказа"})
		return
	}

	BroadcastOrders()

	c.JSON(http.StatusOK, gin.H{
		"message":        "Чек добавлен",
		"receipt_path":   fpath,
		"receipt_paths":  paths,
		"receipts_count": len(paths),
	})
}
