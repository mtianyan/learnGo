package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		// log path, latency, resp code

		s := time.Now()
		c.Next()

		logger.Info(
			"incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
