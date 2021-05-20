package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// add middleware
	r.Use(func(context *gin.Context) {
		startTime := time.Now()
		context.Next() // continue, handle request

		// log: latency, response and response code
		logger.Info(
			"Request: ",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("latency", time.Now().Sub(startTime)),
		)

	}, func(context *gin.Context) {
		context.Set("requestId", rand.Int())

		context.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}

		c.JSON(200, h)
	})
	r.GET("/hi", func(c *gin.Context) {
		c.String(200, "hello there")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
