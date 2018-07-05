package usecase

import (
	"context"
	"time"

	"github.com/hoflish/url-shortener/api/models"
	"github.com/hoflish/url-shortener/api/urlshorten"
	"github.com/teris-io/shortid"
)

// TODO: refactor this code,
const shortBaseURL = "http://192.168.99.100:8080/"

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

func (usu *URLShortenUsecase) Store(c context.Context, us *models.URLShorten) (*models.URLShorten, error) {
	ctx, cancel := context.WithTimeout(c, usu.contextTimeout)
	defer cancel()

	//sid, err := shortid.New(1, shortid.DefaultABC, 2342)

	// TODO: refactor this code to be more safe
	shortID, err := shortid.Generate()
	if err != nil {
		return nil, err
	}

	us.CreatedAt = time.Now()
	us.UpdatedAt = time.Now()
	us.ShortURL = shortBaseURL + shortID

	res, err := usu.urlShortenRepos.Store(ctx, us)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (usu *URLShortenUsecase) Close() {}
