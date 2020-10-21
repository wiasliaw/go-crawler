package config

import (
	"log"

	"github.com/joho/godotenv"
)

const (
	// QueueAmount is the amount of queue
	QueueAmount int = 4096

	// WorkerAmount is the amount of worker
	WorkerAmount int = 1024

	// BenchmarkTimeout is the time to stop the main.go
	BenchmarkTimeout int = 60
)

var (
	// Urls is the start of crawler
	Urls []string = []string{
		"https://mobile.dcard.tw/f/relationship",
		"https://mobile.dcard.tw/f/funny",
		"https://mobile.dcard.tw/f/mood",
		"https://mobile.dcard.tw/f/food",
		"https://mobile.dcard.tw/f/makeup",
		"https://mobile.dcard.tw/f/apple",
		"https://mobile.dcard.tw/f/girl",
		"https://www.dcard.tw/f/talk",
		"https://www.dcard.tw/f/dressup",
		"https://www.dcard.tw/f/pet",
	}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
