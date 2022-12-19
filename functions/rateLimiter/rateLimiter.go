package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)
/*
adds a rate limiter 
example (20, time.Minute, time.Minute)
timeWhenRequestsRefresh / timeWhenOneRequestIsAddedToUser = how many requests per that time for example
60 sec / 10sec  every 10 seconds 1 request is available
20 req
*/
func RateLimiter(requestAmount int, timeWhenOneRequestIsAddedToUser time.Duration, timeWhenRequestsRefresh time.Duration) gin.HandlerFunc {
	return limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP()
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(timeWhenOneRequestIsAddedToUser), requestAmount), timeWhenRequestsRefresh // 20 req / minute | when the requests refill
	}, func(c *gin.Context) {
		// c.JSON(429, gin.H{
			
		// })
		c.AbortWithStatus(429)
	})
}