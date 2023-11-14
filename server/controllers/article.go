package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// article represents data about a blog article.
type article struct {
	Id			uint64 `json:"id"`
	Title		string `json:"title"`
	TopicId		uint64 `json:"topicId"`
	Description	string `json:"description"`
}

// article slice to seed record article data.
var articles = []article{
	{Id: 1, Title: "Coherence", TopicId: 1, Description: "2021-08-25 02:18:45"},
	{Id: 2, Title: "Health", TopicId: 2, Description: "2021-05-09 17:24:45"},
	{Id: 3, Title: "Security", TopicId: 3, Description: "2021-04-12 23:42:45"},
}

func ArticleRoutes(g *gin.RouterGroup) {
	g.GET("/", getArticles)
}

// getArticles responds with the list of all articles as JSON.
func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}