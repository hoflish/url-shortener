package usecase

import (
	"context"
	"time"

	"github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/url"
)

type urlUsecase struct {
	urlRepos       url.UrlRepository
	contextTimeout time.Duration
}

func NewUrlUsecase(u url.UrlRepository, timeout time.Duration) url.UrlRepository {
	return &urlUsecase{
		urlRepos:       u,
		contextTimeout: timeout,
	}
}

func (u *urlUsecase) Fetch(c context.Context, urlId string) (*models.Url, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	item, err := u.urlRepos.Fetch(ctx, urlId)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (u *urlUsecase) Store(c context.Context, url *models.Url) (*models.Url, error) {
	panic("Not Implemented")
}
