package crawler

import (
	"go-crawler/config"
	"log"

	"github.com/gocolly/colly"
)

func (w *Worker) hook() {
	w.collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	w.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		byteLink := []byte(link)
		if !config.ReForum.Match(byteLink) && !config.RePost.Match(byteLink) {
			return
		}
		log.Printf("Link found: %q -> %s\n", e.Text, link)
		w.queue.Push(e.Request.AbsoluteURL(link))
	})
}
