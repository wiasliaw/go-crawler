package main

import (
	"context"

	"go-crawler/config"
	"go-crawler/crawler"
	"go-crawler/db"
	"go-crawler/queue"
	"go-crawler/utils"

	"go.uber.org/fx"
)

func register(
	lifecycle fx.Lifecycle,
	d *db.Database,
	q *queue.Queue,
	ws []*crawler.Worker,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go utils.TimeOut(config.BenchmarkTimeout)
				for _, url := range config.Urls {
					q.Push(url)
				}
				for _, worker := range ws {
					go worker.Run()
				}
				return nil
			},
			OnStop: func(context.Context) error {
				d.Close()
				return nil
			},
		},
	)
}

func main() {
	fx.New(
		fx.Provide(db.Provider),
		fx.Provide(queue.Provider),
		fx.Provide(crawler.Provider),
		fx.Invoke(register),
	).Run()
}
