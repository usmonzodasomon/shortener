package controllers

import (
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/usmonzodasomon/shortener/internal/models"
	"log/slog"
	"net/http"
)

type UserServiceI interface {
	CreateUser(user models.User) (int, error)
}

type UserController struct {
	log     *slog.Logger
	service UserServiceI
}

func NewUserController(log *slog.Logger, service UserServiceI) *UserController {
	return &UserController{
		log:     log,
		service: service,
	}
}

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserRequest
	if err := render.DecodeJSON(r.Body, &user); err != nil {
		c.log.Warn("failed to decode request body", slog.String("error", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "failed to decode request body",
		})
	}

	if err := validator.New().Struct(user); err != nil {
		c.log.Warn("invalid request body", slog.String("error", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request body",
		})
	}

	id, err := c.service.CreateUser(models.User{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		c.log.Error("failed to create user", slog.String("error", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to create user",
		})
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, map[string]any{"user_id": id})
}
