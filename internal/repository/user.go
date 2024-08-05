package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/shortener/internal/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user models.User) (int, error) {
	q := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	var id int
	err := r.db.Get(&id, q, user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	return id, nil
}
