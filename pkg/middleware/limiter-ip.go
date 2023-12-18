package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"lark/pkg/xhttp"
	"sync"
	"time"
)

var (
	ipLimits = make(map[string]*rate.Limiter)
	rwLock   sync.RWMutex
)

func IpLimiter(fillInterval time.Duration, cap int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		rwLock.RLock()
		limiter, exists := ipLimits[ip]
		rwLock.RUnlock()
		if !exists {
			limiter = rate.NewLimiter(rate.Every(fillInterval), cap)
			rwLock.Lock()
			ipLimits[ip] = limiter
			rwLock.Unlock()
		}
		if !limiter.Allow() {
			//ctx.AbortWithStatus(429)
			ctx.Abort()
			xhttp.Error(ctx, xhttp.ERROR_CODE_HTTP_TOO_MANY_REQUESTS, xhttp.ERROR_HTTP_TOO_MANY_REQUESTS)
			return
		}
		ctx.Next()
	}
}
