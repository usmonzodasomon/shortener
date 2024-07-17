package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UrlServiceI interface {
	SaveURL(url string) (string, error)
	GetURL(shortUrl string) (string, error)
	IncrementCount(shortUrl string) error
}

type UrlController struct {
	service UrlServiceI
}

func NewUrlController(service UrlServiceI) *UrlController {
	return &UrlController{
		service: service,
	}
}

type Url struct {
	Url string `json:"url" validate:"required,url"`
}

func (c *UrlController) SaveURL(w http.ResponseWriter, r *http.Request) {
	var url Url
	if err := render.DecodeJSON(r.Body, &url); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if err := validator.New().Struct(url); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}

	shortUrl, err := c.service.SaveURL(url.Url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, map[string]string{"short_url": shortUrl, "url": url.Url})
}

func (c *UrlController) GetURL(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "shortUrl")
	if shortUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "short_url is required",
		})
		return
	}

	url, err := c.service.GetURL(shortUrl)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "short_url not found",
		})
		return
	}

	if err := c.service.IncrementCount(shortUrl); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
