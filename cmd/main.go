package main

import (
	"context"
	"go-crawler/pkg/config"
	"go-crawler/pkg/db"
	"go-crawler/pkg/queue"
	"go-crawler/pkg/worker"
	"log"
	"os"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"go.uber.org/fx"
)

func timeout() {
	select {
	case <-time.After(time.Duration(config.BenchmarkTimeout) * time.Second):
		log.Printf("end")
	}
	os.Exit(0)
}

func provideDB() *db.Database {
	d := db.New()
	d.Init()
	return d
}

func provideQueue() *queue.Queue {
	return queue.New(config.QueueAmount)
}

func provideCache() *lru.Cache {
	c, _ := lru.New(config.CacheSize)
	return c
}

func provideWorkers(
	q *queue.Queue,
	c *lru.Cache,
	d *db.Database,
) []*worker.Worker {
	ws := []*worker.Worker{}
	for i := 0; i < config.WorkerAmount; i++ {
		w := worker.New(
			worker.SetCache(c),
			worker.SetDatabase(d),
			worker.SetQueue(q),
		)
		ws = append(ws, w)
	}
	return ws
}

func register(
	lifecycle fx.Lifecycle,
	q *queue.Queue,
	c *lru.Cache,
	d *db.Database,
	ws []*worker.Worker,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// queue init
			for _, url := range config.Urls {
				q.Push(url)
			}
			// worker run
			for _, w := range ws {
				go w.Run()
			}
			go timeout()
			return nil
		},
		OnStop: func(c context.Context) error {
			d.Close()
			return nil
		},
	})
}

func main() {
	fx.New(
		fx.Provide(provideDB),
		fx.Provide(provideCache),
		fx.Provide(provideQueue),
		fx.Provide(provideWorkers),
		fx.Invoke(register),
	).Run()
}
