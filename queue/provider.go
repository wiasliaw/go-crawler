package queue

import (
	"go-crawler/db"
)

// Provider returns queue dependency
func Provider(d *db.Database) *Queue {
	return New(d)
}
