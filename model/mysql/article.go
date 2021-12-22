package mysql

import (
	"database/sql"
	"errors"
	"github.com/holehole5566/goproject/model"
	"github.com/holehole5566/goproject/repo"
	
)

type mysqlArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) repo.ArticleRepo {
	return &mysqlArticleRepository{db: db}
}

func (m *mysqlArticleRepository) Get(id int) (*model.Article, error) {

	scanID := 0
	scanTitle := ""
	scanContent := ""
	if err := m.db.QueryRow("SELECT id, title,content FROM goblog WHERE id = ?", id).Scan(&scanID, &scanTitle,&scanContent); err != nil {
		return nil, err
	} else if scanID != id {
		return nil, errors.New("scan id and param id are not matched")
	}

	rows, err := m.db.Query(`SELECT id, title ,content FROM goblog 
								where id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var article model.Article
	article.ID = id
	article.Title = scanTitle
	article.Content = scanContent
	return &article, nil
}

func (m *mysqlArticleRepository) Gets() (goblog []*model.Article, err error) {
	rows, err := m.db.Query(`SELECT goblog.id, goblog.title,goblog.content FROM goblog`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var content string
		var tid int
		var title string
		if err := rows.Scan(&tid, &title,&content); err != nil {
			return nil, err
		}
		goblog = append(goblog, &model.Article{ID: tid, Title: title, Content: content})
	}
	return
}

func (m *mysqlArticleRepository) Add(content string, title string) (int, error) {


	tx, err := m.db.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec("INSERT INTO goblog (content, title) VALUES (?, ?)", content, title)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}

func (m *mysqlArticleRepository) Del(id int) error {

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM goblog WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}