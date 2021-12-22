package repo

import (
	"github.com/holehole5566/goproject/model"
)

type ArticleRepo interface {
	Add(content string, title string) (int, error)
	Get(id int) (*model.Article, error)
	Gets() ([]*model.Article, error)
	Del(id int) error
}


var Article ArticleRepo
