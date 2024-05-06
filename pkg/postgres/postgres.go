package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/shortener/internal/config"
)

func GetConnection(cfg config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.PostgresHost, cfg.Database.PostgresPort, cfg.Database.PostgresUser, cfg.Database.PostgresPassword,
		cfg.Database.PostgresDatabase)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseConnection(db *sqlx.DB) error {
	return db.Close()
}
