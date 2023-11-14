package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// article represents data about a blog article.
type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// article slice to seed record article data.
var users = []user{
	{ID: "1", Name: "Coherence", Type: "John Coltrane"},
	{ID: "2", Name: "Health", Type: "Gerry Mulligan"},
	{ID: "3", Name: "Security", Type: "Sarah Vaughan"},
}

func AccountRoutes(g *gin.RouterGroup) {
	g.GET("/all", getUsers)
}

// getArticles responds with the list of all articles as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
