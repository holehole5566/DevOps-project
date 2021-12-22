package article

import (
	"database/sql"
	"strconv"
	"log"
	"github.com/holehole5566/goproject/model"
	"github.com/holehole5566/goproject/repo"
	C "github.com/holehole5566/goproject/pkg/constant"

)

func (srv *Service) GetArticle(param string) (*model.Article, error) {

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Get Article: param id is not a number", param)
		return nil, C.ErrArticleIDNotNumber
	}

	article, err := repo.Article.Get(id)
	if err == sql.ErrNoRows {
		log.Println("Get Article: article id record not found", id)
		return nil, C.ErrArticleNotFound
	} else if err != nil {
		log.Println("Get Article: unknown database error", err.Error())
		return nil, C.ErrDatabase
	}

	return article, nil
}

func (srv *Service) GetAllArticle() ([]*model.Article, error) {
	// TODO: cache
	total, err := repo.Article.Gets()
	if err != nil {
		log.Println("Get Articles: unknown database error, ", err.Error())
		return nil, C.ErrDatabase
	}
	return total, nil
}


func (srv *Service) AddArticle(content string, title string) (int, error) {

	if len(content) == 0 || title == "" {
		return 0, C.ErrArticleAddFormatIncorrect
	}

	id, err := repo.Article.Add(content, title)
	if err != nil {
		log.Println("Add Article: ", err)
		return 0, C.ErrDatabase
	}

	return id, nil
}

func (srv *Service) DelArticle(param string) error {

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal("Del Article: param id is not a number", param)
		return C.ErrArticleIDNotNumber
	}

	if id < 0 {
		return C.ErrArticleDelIDIncorrect
	}

	if article, _ := repo.Article.Get(id); article == nil {
		return C.ErrArticleDelDeleted
	}

	if err := repo.Article.Del(id); err != nil {
		log.Println("Del Article: ", err)
		return C.ErrDatabase
	}

	return nil
}