package crawler

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	timerTime  = time.Second * 1
	urlListKey = "url_list_key"
)

type Crawler struct {
	redisClient redis.Client
}

func New(redisClient redis.Client) *Crawler {
	return &Crawler{
		redisClient: redisClient,
	}
}

func (c *Crawler) Run() {
	c.loop()
}

func (c *Crawler) loop() {
	timeTicker := time.NewTicker(timerTime)
	for {
		select {
		case <-timeTicker.C:
			c.do()
		}
	}
}

func (c *Crawler) do() {
	result, err := c.redisClient.RPop(urlListKey).Result()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("url: ", result)
}
