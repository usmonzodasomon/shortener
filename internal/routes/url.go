package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/shortener/internal/controllers"
	"github.com/usmonzodasomon/shortener/internal/repository"
	"github.com/usmonzodasomon/shortener/internal/service"
	"log/slog"
)

func urlRoutes(r *chi.Mux, db *sqlx.DB, logger *slog.Logger) {
	urlRepo := repository.NewUrlRepository(db)
	urlService := service.NewUrlService(urlRepo)
	urlControllers := controllers.NewUrlController(logger, urlService)

	r.Get("/add", urlControllers.AddURL)
	r.Route("/url", func(r chi.Router) {
		r.Post("/", urlControllers.SaveURL)
	})
	r.Get("/{shortUrl}", urlControllers.GetURL)

}
