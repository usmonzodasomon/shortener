package repository

import "github.com/jmoiron/sqlx"

type UrlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) SaveURL(url string, shortUrl string) error {
	q := "INSERT INTO urls (url, short_url) VALUES ($1, $2)"
	_, err := r.db.Exec(q, url, shortUrl)
	if err != nil {
		return err
	}
	return nil
}

func (r *UrlRepository) GetURL(shortUrl string) (string, error) {
	q := "SELECT url FROM urls WHERE short_url = $1"
	var url string
	err := r.db.Get(&url, q, shortUrl)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (r *UrlRepository) IncrementCount(shortUrl string) error {
	q := "UPDATE urls SET count = count + 1 WHERE short_url = $1"
	_, err := r.db.Exec(q, shortUrl)
	if err != nil {
		return err
	}
	return nil
}
