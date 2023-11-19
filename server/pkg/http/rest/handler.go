package rest

import (
	"net/http"

	"github.com/AkashJam/blog-server/server/pkg/article"
	"github.com/gin-gonic/gin"
)

//Necessary to import local package into another local package
type Config struct {
	R				*gin.Engine
// 	ArticleService	article.Service
}

type Handler struct {
	R				*gin.Engine
// 	ArticleService	article.Service
}

func getHealthy(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"msg": "it works",
	})
}

func NewHandler(c *Config) *Handler {
	h := &Handler{
		R:				c.R,
		// ArticleService:	c.ArticleService,
	}
	g := h.R.Group("/api")
	g.GET("/", getHealthy)
	g.GET("/articles", article.GetArticles)
	return h
}

// func AccountRoutes(g *gin.RouterGroup) {
// 	g.GET("/", article.getArticles)
// }