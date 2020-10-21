package queue

import (
	"go-crawler/config"
	"go-crawler/db"
	"go-crawler/utils"
)

// New returns a Queue Reference
func New(q *db.Database) *Queue {
	instance := &Queue{
		jobs: make(chan string, config.QueueAmount),
		db:   q,
	}
	return instance
}

// Queue is a struct
type Queue struct {
	jobs chan string
	db   *db.Database
}

// Push an url into queue
func (q *Queue) Push(url string) {
	select {
	case q.jobs <- url:
		{
			hash := utils.MD5Hash(url)
			q.db.Insert(hash)
		}
	default:
	}
}

// Pop an url to worker
func (q *Queue) Pop() string {
	url := <-q.jobs
	return url
}
