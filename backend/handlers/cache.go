package handlers

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

// dataVersion is bumped whenever orders or the product catalogue change (see
// BroadcastOrders / BroadcastProducts). Cached analytics responses record the version
// they were computed at and are treated as stale the moment it moves, so a cache hit can
// never serve numbers from before the latest sale.
var dataVersion int64

// bumpDataVersion invalidates every cached read. Called from the broadcast helpers, which
// already fire on every order/product mutation.
func bumpDataVersion() { atomic.AddInt64(&dataVersion, 1) }

type cacheEntry struct {
	body        []byte
	contentType string
	version     int64
	expiry      time.Time
}

var (
	respCacheMu sync.Mutex
	respCache   = make(map[string]cacheEntry)
)

// cacheKey is scoped by method, exact path (incl. :id params), query and the caller's
// identity, so one worker/doctor/admin can never receive another's cached analytics.
func cacheKey(c *gin.Context) string {
	var id string
	for _, k := range []string{"workerID", "doctorID", "userID", "adminID"} {
		if v, ok := c.Get(k); ok {
			id += fmt.Sprintf("%s=%v;", k, v)
		}
	}
	return c.Request.Method + " " + c.Request.URL.Path + "?" + c.Request.URL.RawQuery + "|" + id
}

// cacheCapturer buffers the response body so a successful read can be cached.
type cacheCapturer struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (w *cacheCapturer) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *cacheCapturer) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// CacheGET caches successful (200) GET responses for up to ttl, keyed per-identity and
// invalidated immediately when orders/products change. This makes repeated and concurrent
// dashboard fetches (multiple admins, tab switches, re-renders) free instead of
// re-scanning the orders table — without altering a single computed value, since a real
// data change bumps the version and forces a fresh recompute.
func CacheGET(ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := cacheKey(c)
		ver := atomic.LoadInt64(&dataVersion)

		respCacheMu.Lock()
		e, ok := respCache[key]
		respCacheMu.Unlock()
		if ok && e.version == ver && time.Now().Before(e.expiry) {
			c.Data(200, e.contentType, e.body)
			c.Abort()
			return
		}

		cap := &cacheCapturer{ResponseWriter: c.Writer, body: &bytes.Buffer{}, status: 200}
		c.Writer = cap
		c.Next()

		if cap.status == 200 {
			respCacheMu.Lock()
			// Bound the map so distinct queries can't leak memory forever.
			if len(respCache) > 500 {
				respCache = make(map[string]cacheEntry)
			}
			respCache[key] = cacheEntry{
				body:        cap.body.Bytes(),
				contentType: cap.Header().Get("Content-Type"),
				version:     ver,
				expiry:      time.Now().Add(ttl),
			}
			respCacheMu.Unlock()
		}
	}
}
