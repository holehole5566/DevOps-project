package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/holehole5566/goproject/router/api/article"
)

func methodNotAllowed(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
}

func InitRouters() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile("./web/dist", true)))


	r.HandleMethodNotAllowed = true

	r.NoMethod(methodNotAllowed)

	articleAPI := r.Group("/api/article")

	articleAPI.GET("/:id", article.GetArticle)

	articleAPI.POST("/", article.AddArticle)

	articleAPI.GET("/", article.GetAllArticle)

	articleAPI.DELETE("/:id", article.DelArticle)

	r.NoRoute(static.Serve("/", static.LocalFile("./web/dist", true)))

	return r
}