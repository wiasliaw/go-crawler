package queue

// Provider returns queue dependency
func Provider(amount int) *Queue {
	return New(amount)
}
