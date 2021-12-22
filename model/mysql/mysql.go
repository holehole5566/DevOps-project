package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/holehole5566/goproject/repo"
)

func Setup() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"hole",
		"123",
		"db",
		"goblog"))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	repo.Article = NewArticleRepository(db)

}

