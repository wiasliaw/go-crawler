package worker

import (
	"go-crawler/pkg/db"
	"go-crawler/pkg/queue"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	lru "github.com/hashicorp/golang-lru"
)

// Worker struct
type Worker struct {
	url       string
	idle      chan bool
	queue     *queue.Queue
	collector *colly.Collector
	cache     *lru.Cache
	db        *db.Database
}

// New Initialize and return Worker's reference
func New(opts ...Option) *Worker {
	instance := &Worker{
		idle:      make(chan bool, 1),
		collector: collySetup(),
	}
	for _, opt := range opts {
		opt.set(instance)
	}
	instance.hook()
	return instance
}

// Run worker run
func (w *Worker) Run() {
	w.idle <- true
	for {
		select {
		case <-w.idle:
			go w.Visit()
		}
	}
}

// Visit worker visit
func (w *Worker) Visit() {
	w.url = w.queue.Pop()
	// check if contain in cache
	isContained, _ := w.cache.ContainsOrAdd(
		hashMD5(w.url),
		w.url,
	)
	if !isContained {
		w.collector.Visit(w.url)
	}
	w.idle <- true
}

// hook
func (w *Worker) hook() {
	// onRequest event - log
	w.collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
		go w.db.Insert(hashMD5(r.URL.String()))
	})

	// onError event - log error
	w.collector.OnError(func(r *colly.Response, err error) {
		log.Println("ERROR", w.collector.ID, r.Request.URL, err)
	})

	// onResponse event - label type
	w.collector.OnResponse(func(r *colly.Response) {
		// log.Println(contentTypeFilter(r.Headers.Get("Content-Type")))
	})

	// onHTML event - parser tag to url queue
	w.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		absLink := e.Request.AbsoluteURL(link)

		// remove query
		url, err := queryFilter(absLink)
		if err != nil {
			return
		}

		// log.Println(url)
		w.queue.Push(url)
	})

	// final event - db insert
	w.collector.OnScraped(func(r *colly.Response) {
		// w.db.Insert(hashMD5(r.Request.URL.String()))
	})
}

// collySetup initialize colly instance
func collySetup() *colly.Collector {
	// colly collector
	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)
	// extensions and setup
	// c.SetRequestTimeout(time.Duration(3) * time.Second)
	extensions.Referer(c)
	extensions.RandomUserAgent(c)
	return c
}
