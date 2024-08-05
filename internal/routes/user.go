package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/usmonzodasomon/shortener/internal/controllers"
	"github.com/usmonzodasomon/shortener/internal/repository"
	"github.com/usmonzodasomon/shortener/internal/service"
	"log/slog"
)

func userRoutes(r *chi.Mux, db *sqlx.DB, logger *slog.Logger) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userControllers := controllers.NewUserController(logger, userService)

	r.Post("/sign-up", userControllers.CreateUser)
}
