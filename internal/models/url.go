package models

type Url struct {
	ID       int    `db:"id"`
	Url      string `db:"url"`
	ShortUrl string `db:"short_url"`
	Count    int    `db:"count"`
}
