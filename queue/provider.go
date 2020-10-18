package queue

// Provider returns queue dependency
func Provider() *Queue {
	return New()
}
