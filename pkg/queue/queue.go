package queue

// Queue type for queue
type Queue struct {
	urls chan string
}

// New initialize and return Queue's reference
func New(amount int) *Queue {
	instance := &Queue{
		urls: make(chan string, amount),
	}
	return instance
}

// Push an url into queue
func (q *Queue) Push(url string) {
	select {
	case q.urls <- url:
	default:
	}
}

// Pop an url to worker
func (q *Queue) Pop() string {
	url := <-q.urls
	return url
}
