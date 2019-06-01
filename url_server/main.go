package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const (
	urlListKey = "url_list_key"
)

type Req struct {
	URL string `json:"url"`
}

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password",
	})
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var req Req
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		redisClient.RPush(urlListKey, req.URL)
		c.JSON(200, gin.H{})
	})

	r.Run()
}
