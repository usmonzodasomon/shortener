package service

import "github.com/usmonzodasomon/shortener/pkg/random"

type UrlRepositoryI interface {
	SaveURL(url string, shortUrl string) error
	GetURL(shortUrl string) (string, error)
	IncrementCount(shortUrl string) error
}

type UrlService struct {
	repo UrlRepositoryI
}

func NewUrlService(repo UrlRepositoryI) *UrlService {
	return &UrlService{
		repo: repo,
	}
}

const (
	shortUrlLength = 6
)

func (s *UrlService) SaveURL(url string) (string, error) {
	shortUrl := random.NewString(shortUrlLength)
	if err := s.repo.SaveURL(url, shortUrl); err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (s *UrlService) GetURL(shortUrl string) (string, error) {
	url, err := s.repo.GetURL(shortUrl)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *UrlService) IncrementCount(shortUrl string) error {
	if err := s.repo.IncrementCount(shortUrl); err != nil {
		return err
	}
	return nil
}
