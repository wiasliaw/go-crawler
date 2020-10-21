package crawler

import (
	"go-crawler/db"
	"go-crawler/queue"

	"github.com/gocolly/colly"
)

// New returns worker dependency
func New(q *queue.Queue) *Worker {
	instance := &Worker{
		queue: q,
		idle:  make(chan bool, 1),
		collector: colly.NewCollector(
			colly.AllowedDomains("mobile.dcard.tw", "www.dcard.tw"),
		),
	}
	instance.hook()
	return instance
}

// Worker is a struct
type Worker struct {
	url       string
	idle      chan bool
	collector *colly.Collector
	db        *db.Database
	queue     *queue.Queue
}

// Run is a function run worker
func (w *Worker) Run() {
	w.idle <- true
	for {
		select {
		case <-w.idle:
			go w.Visit()
		}
	}
}

// Visit is a function to Visit url
func (w *Worker) Visit() {
	w.url = w.queue.Pop()
	w.collector.Visit(w.url)
	w.idle <- true
}
