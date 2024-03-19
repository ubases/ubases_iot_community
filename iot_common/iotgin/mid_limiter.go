package iotgin

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"golang.org/x/time/rate"
)

type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

var (
	RateLimiter *IPRateLimiter
)

func SetupIPRateLimiter(qps int) error {
	RateLimiter = &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   1,
		b:   qps,
	}
	return nil
}

// 添加一个ip到map
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

//通过ip得到limiter
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}

//IP限流器
func LimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := RateLimiter.GetLimiter(c.ClientIP())
		if !limiter.Allow() {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"code": http.StatusServiceUnavailable,
				"msg":  "访问频率过高,请稍后访问",
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

//服务限流
func Limiter(maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second*1), maxBurstSize)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code": http.StatusServiceUnavailable,
			"msg":  "服务超负荷运行,请稍后访问",
			"data": nil,
		})
		c.Abort()
	}
}
