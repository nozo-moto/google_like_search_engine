package crawler

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/go-redis/redis"
	"github.com/sclevine/agouti"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	timerTime  = time.Second * 1
	urlListKey = "url_list_key"
)

type Crawler struct {
	redisClient redis.Client
	driver      *agouti.WebDriver
	db          *sqlx.DB
}

func New(redisClient redis.Client) (*Crawler, error) {
	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
		}),
		agouti.Debug,
	)
	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf("mysql:password@tcp(localhost:3306)/google_like_search_engine?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo"),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("chrome driver init and db connected")
	return &Crawler{
		redisClient: redisClient,
		driver:      driver,
		db:          db,
	}, err
}

func (c *Crawler) Run() {
	c.loop()
}

func (c *Crawler) loop() {
	err := c.driver.Start()
	if err != nil {
		log.Printf("Failed to start driver: %v", err)
	}

	timeTicker := time.NewTicker(timerTime)
	for {
		select {
		case <-timeTicker.C:
			c.do()
		}
	}
}

func (c *Crawler) do() {
	// get url from redis
	gotUrl, err := c.redisClient.RPop(urlListKey).Result()
	if err != nil {
		return
	}

	page, err := c.scraiping(gotUrl)
	if err != nil {
		panic(err)
	}

	page, err = c.insertPage(page)
	if err != nil {
		panic(err)
	}

	fmt.Println("inserted data", page.ID, page.Title)
}

func (c *Crawler) scraiping(gotUrl string) (*Page, error) {
	// chromium setup
	page, err := c.driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		return nil, err
	}
	err = page.Navigate(gotUrl)
	if err != nil {
		return nil, err
	}

	html, err := page.HTML()
	if err != nil {
		return nil, err
	}

	title, err := page.Title()
	if err != nil {
		return nil, err
	}

	u, err := page.URL()
	if err != nil {
		return nil, err
	}

	u, err = url.QueryUnescape(u)
	if err != nil {
		return nil, err
	}

	p := &Page{
		Title: title,
		Url:   u,
		Html:  html,
	}

	return p, err
}

func (c *Crawler) insertPage(page *Page) (*Page, error) {
	fmt.Println(page.Title, page.Url)
	result, err := c.db.Exec(
		`INSERT INTO page(page_title, page_url, page_html) VALUES(?, ?, ?)`,
		page.Title, page.Url, page.Html,
	)
	if err != nil {
		return page, err
	}
	page.ID, err = result.LastInsertId()
	return page, err
}
