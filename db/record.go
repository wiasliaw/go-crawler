package db

// Record is a struct for query result
type Record struct {
	ID      int    `db:"url_id" json:"id"`
	URLHash string `db:"url_hash" json:"url_hash"`
}
