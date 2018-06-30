package usecase

import (
	"context"
	"time"

	"github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/urlshorten"
)

type URLShortenUsecase struct {
	urlShortenRepos urlshorten.URLShortenRepos
	contextTimeout  time.Duration
}

func NewURLShortenUsecase(us urlshorten.URLShortenRepos, timeout time.Duration) urlshorten.URLShortenRepos {
	return &URLShortenUsecase{
		urlShortenRepos: us,
		contextTimeout:  timeout,
	}
}

// Fetch implements business loginc of ShortURL fetching
func (usu *URLShortenUsecase) Fetch(c context.Context, shortURL string) (*models.URLShorten, error) {
	ctx, cancel := context.WithTimeout(c, usu.contextTimeout)
	defer cancel()

	item, err := usu.urlShortenRepos.Fetch(ctx, shortURL)

	if err != nil {
		return nil, err
	}

	return item, nil
}

/*func (usu *URLShortenUsecase) Store(c context.Context, urlShorten *models.URLShorten) (*models.URLShorten, error) {
	panic("Not Implemented")
}*/
