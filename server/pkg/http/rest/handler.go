package rest

import (
	. "github.com/AkashJam/blog-server/server/pkg/adding"
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(g *gin.RouterGroup) {
	g.GET("/", getArticles)
}