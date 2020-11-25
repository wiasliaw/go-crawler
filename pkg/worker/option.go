package worker

import (
	"go-crawler/pkg/db"
	"go-crawler/pkg/queue"

	lru "github.com/hashicorp/golang-lru"
)

//Option abstraction for config
type Option interface {
	set(*Worker)
}

// OptsFunc Option's type implement
type OptsFunc func(*Worker)

func (f OptsFunc) set(w *Worker) { f(w) }

// SetDatabase setup options for db
func SetDatabase(d *db.Database) OptsFunc {
	return OptsFunc(func(w *Worker) {
		w.db = d
	})
}

// SetQueue setup options for queue
func SetQueue(q *queue.Queue) OptsFunc {
	return OptsFunc(func(w *Worker) {
		w.queue = q
	})
}

//SetCache setup options for cache
func SetCache(l *lru.Cache) OptsFunc {
	return OptsFunc(func(w *Worker) {
		w.cache = l
	})
}
