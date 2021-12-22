package article

import (
	"net/http"
	"github.com/holehole5566/goproject/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/holehole5566/goproject/service"
)

type Article struct {
	title string  `json:"title"`
	subject    string `json:"content"`
	date       string    `json:"date"`
}

func GetArticle(c *gin.Context) {

	appG := app.Gin{C: c}

	switch tour, err := service.article.GetArticle(c.Param("id")); err {

	case C.ErrTourIDNotNumber:
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error(), nil)

	case C.ErrTourNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_TOUR_NO_RECORD, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_TOUR_FAIL, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, tour)

	default:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)
	}
}

func GetAllArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	switch total, err := service.article.GetAllArticle(); err {

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, total)

	}
}


func AddArticle(c *gin.Context) {

	appG := app.Gin{C: c}

	var t article

	if err := c.BindJSON(&t); err != nil {
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error(), nil)
		return
	}

	switch tourID, err := service.Game.AddTour(t.Collects, t.Title); err {

	case C.ErrTourAddFormatIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_TOUR_FORMAT_INCORRECT, err.Error(), nil)

	case C.ErrTourAddCollectsRecordNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_TOUR_NO_COLLECTS_RECORD, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, tourID)

	}
}

func DelArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	switch err := service.article.DelArticle(c.Param("id")); err {

	case C.ErrTourDelIDIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_DEL_TOUR_ID_INCORRECT, err.Error(), nil)

	case C.ErrTourDelDeleted:
		appG.Response(http.StatusGone, C.ERROR_DEL_TOUR_DELETED, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, nil)

	}
}