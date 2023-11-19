package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) getArticles(c *gin.Context) {
	item, _ := h.ArticleService.GetArticles()
	c.IndentedJSON(http.StatusOK, item)
}