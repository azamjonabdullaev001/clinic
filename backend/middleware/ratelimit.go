package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// visitor is a per-IP token bucket.
type visitor struct {
	tokens   float64
	lastSeen time.Time
}

type rateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	rate     float64 // tokens refilled per second
	burst    float64 // bucket capacity
}

// RateLimit returns a middleware that throttles each client IP to roughly `rps`
// requests per second with room for short bursts. It is a lightweight in-process
// guard (no Redis needed for a single backend instance) that protects the API and
// database from accidental floods and abusive clients. A background sweeper evicts
// idle IPs so the map cannot grow without bound.
func RateLimit(rps float64, burst float64) gin.HandlerFunc {
	rl := &rateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rps,
		burst:    burst,
	}

	go func() {
		for range time.Tick(time.Minute) {
			rl.mu.Lock()
			for ip, v := range rl.visitors {
				if time.Since(v.lastSeen) > 3*time.Minute {
					delete(rl.visitors, ip)
				}
			}
			rl.mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		rl.mu.Lock()
		v, ok := rl.visitors[ip]
		now := time.Now()
		if !ok {
			v = &visitor{tokens: rl.burst, lastSeen: now}
			rl.visitors[ip] = v
		} else {
			// Refill since the last request, capped at burst.
			v.tokens += now.Sub(v.lastSeen).Seconds() * rl.rate
			if v.tokens > rl.burst {
				v.tokens = rl.burst
			}
			v.lastSeen = now
		}

		if v.tokens < 1 {
			rl.mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Слишком много запросов. Попробуйте через несколько секунд.",
			})
			return
		}
		v.tokens--
		rl.mu.Unlock()

		c.Next()
	}
}
