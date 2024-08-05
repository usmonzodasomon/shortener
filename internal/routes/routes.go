package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

func SetUpRoutes(r *chi.Mux, db *sqlx.DB, logger *slog.Logger) {
	urlRoutes(r, db, logger)
	userRoutes(r, db, logger)
}
