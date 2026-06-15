package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "" {
			return true // same-origin requests have no Origin header
		}
		return origin == "http://"+r.Host || origin == "https://"+r.Host
	},
}

type stockHubT struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

var stockHub = &stockHubT{clients: make(map[*websocket.Conn]bool)}

func (h *stockHubT) add(c *websocket.Conn) {
	h.mu.Lock()
	h.clients[c] = true
	h.mu.Unlock()
}

func (h *stockHubT) remove(c *websocket.Conn) {
	h.mu.Lock()
	delete(h.clients, c)
	h.mu.Unlock()
	c.Close()
}

func (h *stockHubT) broadcast(msg interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for c := range h.clients {
		if err := c.WriteJSON(msg); err != nil {
			c.Close()
			delete(h.clients, c)
		}
	}
}

// StockWebSocket upgrades the connection and keeps it open to push stock updates.
func StockWebSocket(c *gin.Context) {
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	stockHub.add(conn)
	defer stockHub.remove(conn)
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}

// BroadcastStock pushes a product's current stock quantity to all clients.
func BroadcastStock(productID uint, quantity int) {
	stockHub.broadcast(gin.H{"type": "stock", "product_id": productID, "quantity": quantity})
}

// broadcastProductStock reads the product's current quantity and broadcasts it.
func broadcastProductStock(productID uint) {
	var p models.Product
	if database.DB.Select("id", "stock_quantity").First(&p, productID).Error == nil {
		BroadcastStock(p.ID, p.StockQuantity)
	}
}

// BroadcastOrders signals every client that orders changed, so panels can refresh
// their order lists and analytics in real time.
func BroadcastOrders() {
	stockHub.broadcast(gin.H{"type": "orders"})
}

// BroadcastProducts signals every client that the product catalog changed (a product was
// added / edited / deleted), so all panels refresh their product lists in real time.
func BroadcastProducts() {
	stockHub.broadcast(gin.H{"type": "products"})
}
