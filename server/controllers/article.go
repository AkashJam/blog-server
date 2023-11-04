package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// article represents data about a blog article.
type article struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Date   string `json:"date"`
}

// article slice to seed record article data.
var articles = []article{
	{ID: "1", Title: "Coherence", Author: "John Coltrane", Date: "2021-08-25 02:18:45"},
	{ID: "2", Title: "Health", Author: "Gerry Mulligan", Date: "2021-05-09 17:24:45"},
	{ID: "3", Title: "Security", Author: "Sarah Vaughan", Date: "2021-04-12 23:42:45"},
}

func ArticleRoutes(g *gin.RouterGroup) {
	g.GET("/all", getArticles)
}

// getArticles responds with the list of all articles as JSON.
func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}
