package service

import (
	"github.com/holehole5566/goproject/model"

	"github.com/holehole5566/goproject/service/article"
)

type ArticleService interface {

	AddArticle(content string, title string) (int, error)
	GetAllArticle() (int, error)
	GetArticle(param string) (*model.Article, error)
	DelArticle(param string) error

}


var Article ArticleService


func init() {
	Article = article.NewArticleService()
}