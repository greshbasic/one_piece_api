package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]

	if !exists {
		limiter = rate.NewLimiter(1, 5)
		visitors[ip] = limiter
	}

	return limiter
}

func rateLimitHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		ip := context.ClientIP()
		limiter := getVisitor(ip)

		if !limiter.Allow() {
			context.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}

		context.Next()
	}
}
