package repository

import (
	"github.com/hawari17/hello-go/article"
)

type ArticleRepository interface {
	Store(a *article.Article) error
	FindByUrl(url string) (*article.Article, error)
}