package mysql

import (
	"database/sql"
	"errors"
	"time"

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
	if err := m.db.QueryRow("SELECT id, title FROM goblog WHERE id = ? AND deleted_at IS NULL", id).Scan(&scanID, &scanTitle); err != nil {
		return nil, err
	} else if scanID != id {
		return nil, errors.New("scan id and param id are not matched")
	}

	rows, err := m.db.Query(`SELECT collects.title, collects.id FROM collects 
								INNER JOIN tour_collects ON collects.id = tour_collects.collect_id 
								AND tour_collects.tour_id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tour model.Article
	tour.ID = id
	tour.Title = scanTitle
	for rows.Next() {
		title, id := "", 0
		rows.Scan(&title, &id)
	}
	return &tour, nil
}

func (m *mysqlArticleRepository) GetTotal() (int, error) {
	count := 0
	err := m.db.QueryRow("SELECT COUNT(*) FROM goblog WHERE deleted_at IS NULL").Scan(&count)
	return count, err
}

func (m *mysqlArticleRepository) Gets() (goblog []*model.Article, err error) {
	rows, err := m.db.Query(`SELECT goblog.id, goblog.title FROM goblog WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tid int
		var title string
		if err := rows.Scan(&tid, &title); err != nil {
			return nil, err
		}
		goblog = append(goblog, &model.Article{ID: tid, Title: title})
	}
	return
}

func (m *mysqlArticleRepository) Add(content string, title string) (int, error) {
	cur := time.Now()
	tx, err := m.db.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec("INSERT INTO goblog (created_at, updated_at, deleted_at, title) VALUES (?, ?, ?, ?)", cur, cur, sql.NullTime{}, title)
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

	_, err = tx.Exec("DELETE FROM tour_collects WHERE tour_id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM goblog WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}