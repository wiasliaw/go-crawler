package queue

import (
	"go-crawler/config"
)

// New returns a Queue Reference
func New() *Queue {
	instance := &Queue{
		jobs: make(chan string, config.QueueAmount),
	}
	return instance
}

// Queue is a struct
type Queue struct {
	jobs chan string
}

// Push an url into queue
func (q *Queue) Push(url string) {
	select {
	case q.jobs <- url:
	default:
	}
}

// Pop an url to worker
func (q *Queue) Pop() string {
	url := <-q.jobs
	return url
}
