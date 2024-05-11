package service

import "github.com/usmonzodasomon/shortener/internal/repository"

type Service struct {
	UrlService *UrlService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UrlService: NewUrlService(repo.UrlRepository),
	}
}
