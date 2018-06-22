package usecase

import (
	"context"
	"time"

	"github.com/hoflish/url-shortener/api/models"

	"github.com/hoflish/url-shortener/api/urlshorten"
)

type urlshortenUsecase struct {
	urlshortenRepos urlshorten.URLShortenRepository
	contextTimeout  time.Duration
}

func NewURLShortenUsecase(u urlshorten.URLShortenRepository, timeout time.Duration) urlshorten.URLShortenUsecase {
	return &urlshortenUsecase{
		urlshortenRepos: u,
		contextTimeout:  timeout,
	}
}

func (u *urlshortenUsecase) Fetch(ctx context.Context, urlCode string) (*models.URLShorten, error) {
	panic("Not Implemented")
}

func (u *urlshortenUsecase) Store(ctx context.Context, urlShorten *models.URLShorten) (string, error) {
	panic("Not Implemented")
}

func (u *urlshortenUsecase) Close() {
	panic("Not implemented")
}
