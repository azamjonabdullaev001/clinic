package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// paginate reads ?limit and ?offset from the request and returns sane,
// bounded values. Every list endpoint MUST use this so a growing orders
// table can never be loaded into memory in full — the single biggest
// scalability hazard in the old code (Find(&orders) with no LIMIT).
//
//	def  – default page size when the client sends no ?limit
//	max  – hard upper bound the client can never exceed
func paginate(c *gin.Context, def, max int) (limit, offset int) {
	limit = def
	if l, err := strconv.Atoi(c.Query("limit")); err == nil && l > 0 {
		limit = l
	}
	if limit > max {
		limit = max
	}
	if o, err := strconv.Atoi(c.Query("offset")); err == nil && o > 0 {
		offset = o
	}
	return limit, offset
}
