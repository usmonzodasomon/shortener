package service

import "github.com/usmonzodasomon/shortener/pkg/random"

type UrlRepositoryI interface {
	SaveURL(url string, shortUrl string) error
	GetURL(shortUrl string) (string, error)
	IncrementCount(shortUrl string) error
}

type UrlService struct {
	UrlRepository UrlRepositoryI
}

func NewUrlService(repo UrlRepositoryI) *UrlService {
	return &UrlService{
		UrlRepository: repo,
	}
}

const (
	shortUrlLength = 6
)

func (s *UrlService) SaveURL(url string) error {
	shortUrl := random.NewString(shortUrlLength)
	if err := s.UrlRepository.SaveURL(url, shortUrl); err != nil {
		return err
	}
	return nil
}
