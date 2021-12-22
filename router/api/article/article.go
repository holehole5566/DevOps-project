package article

import (
	"net/http"
	"github.com/holehole5566/goproject/pkg/app"
	C "github.com/holehole5566/goproject/pkg/constant"
	"github.com/gin-gonic/gin"
	"github.com/holehole5566/goproject/service"
)

type Article struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	Content	 string 	`json:"content"`
}

func GetArticle(c *gin.Context) {

	appG := app.Gin{C: c}

	switch article, err := service.Article.GetArticle(c.Param("id")); err {

	case C.ErrArticleIDNotNumber:
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error(), nil)

	case C.ErrArticleNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_ARTICLE_NO_RECORD, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_ARTICLE_FAIL, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, article)

	default:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)
	}
}

func GetAllArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	switch total, err := service.Article.GetAllArticle(); err {

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, total)

	}
}


func AddArticle(c *gin.Context) {

	appG := app.Gin{C: c}

	var t Article

	if err := c.BindJSON(&t); err != nil {
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error(), nil)
		return
	}

	switch articleID, err := service.Article.AddArticle(t.Content, t.Title); err {

	case C.ErrArticleAddFormatIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_ARTICLE_FORMAT_INCORRECT, err.Error(), nil)

	case C.ErrArticleAddCollectsRecordNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_ARTICLE_NO_COLLECTS_RECORD, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, articleID)

	}
}

func DelArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	switch err := service.Article.DelArticle(c.Param("id")); err {

	case C.ErrArticleDelIDIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_DEL_ARTICLE_ID_INCORRECT, err.Error(), nil)

	case C.ErrArticleDelDeleted:
		appG.Response(http.StatusGone, C.ERROR_DEL_ARTICLE_DELETED, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, nil)

	}
}