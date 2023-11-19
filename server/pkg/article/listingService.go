package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// article slice to seed record article data.
var articles = []article{
	{Id: 1, Title: "Coherence", TopicId: 1, Description: "2021-08-25 02:18:45"},
	{Id: 2, Title: "Health", TopicId: 2, Description: "2021-05-09 17:24:45"},
	{Id: 3, Title: "Security", TopicId: 3, Description: "2021-04-12 23:42:45"},
}

// getArticles responds with the list of all articles as JSON. the *service is to implement the method
func GetArticles(c *gin.Context) {
	// try {
		c.IndentedJSON(http.StatusOK, articles)
		// return articles, nil
	// } catch (error) {
	// 	return []article, error
	// }

}