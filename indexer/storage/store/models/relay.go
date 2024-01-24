package models

type Relay struct {
	ID  uint64 `db:"id"`
	URL string `db:"url"`
}
