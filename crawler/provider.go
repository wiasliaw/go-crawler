package crawler

import (
	"go-crawler/config"
	"go-crawler/queue"
)

// Provider returns worker dependency
func Provider(q *queue.Queue) []*Worker {
	workers := []*Worker{}
	for i := 0; i < config.WorkerAmount; i++ {
		w := New(q)
		workers = append(workers, w)
	}
	return workers
}
