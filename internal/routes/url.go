package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/shortener/internal/controllers"
	"github.com/usmonzodasomon/shortener/internal/repository"
	"github.com/usmonzodasomon/shortener/internal/service"
)

func urlRoutes(r *chi.Mux, db *sqlx.DB) {
	urlRepo := repository.NewUrlRepository(db)
	urlService := service.NewUrlService(urlRepo)
	urlControllers := controllers.NewUrlController(urlService)

	r.Route("/url", func(r chi.Router) {
		r.Post("/", urlControllers.SaveURL)
	})
	r.Get("/{shortUrl}", urlControllers.GetURL)

}
