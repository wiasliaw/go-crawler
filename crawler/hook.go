package crawler

import (
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

var (
	reForum = regexp.MustCompile(`^/f/[a-z]+(\?latest=true)*$`)
	rePost  = regexp.MustCompile(`^/f/[\d\S]+/p/\d+$`)
)

func (w *Worker) hook() {
	w.collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	w.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		byteLink := []byte(link)
		if !reForum.Match(byteLink) && !rePost.Match(byteLink) {
			return
		}
		log.Printf("Link found: %q -> %s\n", e.Text, link)
		w.queue.Push(e.Request.AbsoluteURL(link))
	})
}
