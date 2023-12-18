package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

var limiters = rate.NewLimiter(4, 1)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试获取一个令牌，如果没有令牌则返回限流错误
		if !limiters.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}
