package config

import (
	"log"
	"regexp"

	"github.com/joho/godotenv"
)

const (
	// WorkerAmount is the amount of worker
	WorkerAmount int = 512

	// BenchmarkTimeout is the time to stop the main.go
	BenchmarkTimeout int = 180
)

var (
	// Urls is the start of crawler
	Urls []string = []string{
		"https://mobile.dcard.tw",
	}

	// ReForum is regexp for forum
	ReForum = regexp.MustCompile(`^/f/[a-z]+(\?latest=true)*$`)

	// RePost is regexp for post
	RePost = regexp.MustCompile(`^/f/[\d\S]+/p/\d+$`)
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
