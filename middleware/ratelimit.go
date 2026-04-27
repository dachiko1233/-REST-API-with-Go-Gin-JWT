package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	limitergin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiter() gin.HandlerFunc {
	// 5 requests per minute per IP
	rate, _ := limiter.NewRateFromFormatted("5-M")
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	return limitergin.NewMiddleware(instance)
}
