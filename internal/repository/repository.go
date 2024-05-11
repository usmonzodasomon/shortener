package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	UrlRepository *UrlRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UrlRepository: NewUrlRepository(db),
	}
}
