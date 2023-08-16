package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/khaosles/gtools2/components/g/result"
)

/*
   @File: rate_limit.go
   @Author: khaosles
   @Time: 2023/6/15 00:16
   @Desc:
*/

// RateLimiter 定义限流器结构体
type RateLimiter struct {
	limiter *ratelimit.Bucket
}

// NewRateLimiter 创建并返回一个限流器实例
func NewRateLimiter(fillInterval time.Duration, capacity int64) *RateLimiter {
	return &RateLimiter{
		limiter: ratelimit.NewBucket(fillInterval, capacity),
	}
}

// RateLimiterMiddleware 限流中间件
func (r *RateLimiter) RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if r.limiter.TakeAvailable(1) < 1 {
			c.JSON(http.StatusTooManyRequests, result.NewJsonResult().No(result.LIMIT_ERROR))
			c.Abort()
			return
		}
		c.Next()
	}
}
