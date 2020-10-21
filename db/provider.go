package db

// Provider is a function for providing dependency
func Provider() *Database {
	d := New()
	d.Init()
	return d
}
