package rest

import (
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(g *gin.RouterGroup) {
	g.GET("/", adding.getArticles)
}