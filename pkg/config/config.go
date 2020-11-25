package config

import (
	"log"

	"github.com/joho/godotenv"
)

const (
	// QueueAmount is the amount of queue
	QueueAmount int = 4096

	// WorkerAmount is the amount of worker
	WorkerAmount int = 256

	// BenchmarkTimeout is the time to stop the main.go
	BenchmarkTimeout int = 60

	// CacheSize is the size for lru cache
	CacheSize int = 2048
)

var (
	// Urls is the start of crawler
	Urls []string = []string{

		// Search-engine
		"https://tw.yahoo.com/",
		"https://www.yam.com/",

		// E-commerce
		"https://shopee.tw/",
		"https://24h.pchome.com.tw/",
		"https://www.momoshop.com.tw/main/Main.jsp",
		// "https://www.ruten.com.tw/",
		// "https://www.books.com.tw/",
		"https://www.rakuten.com.tw/",
		"https://www.buy123.com.tw/",
		"https://www.pcone.com.tw/",
		"https://www.etmall.com.tw/",
		"https://tw.carousell.com/",
		"https://www.ebay.com/",
		"https://www.amazon.com/",

		// Media
		"https://udn.com/news/index",
		"https://www.ltn.com.tw/",
		"https://www.chinatimes.com/?chdtv",
		// "https://news.ebc.net.tw/",
		// "https://www.setn.com/",
		"https://www.nownews.com/",

		// Forum
		"https://www.pixnet.net/",
		"https://www.gamer.com.tw/",
		"https://www.mobile01.com/",
		"https://www.dcard.tw/f",
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
