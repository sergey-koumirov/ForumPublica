package ctrl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchResult struct {
	Id   int64
	Name string
}

type SearchResults []SearchResult

func AppSearchItemType(c *gin.Context) {
	temp := make(SearchResults, 0)

	temp = append(temp, SearchResult{Id: 1, Name: "T1"})
	temp = append(temp, SearchResult{Id: 2, Name: "T2"})
	temp = append(temp, SearchResult{Id: 3, Name: "T3"})
	temp = append(temp, SearchResult{Id: 4, Name: "T4"})
	temp = append(temp, SearchResult{Id: 5, Name: "T5"})

	c.JSON(http.StatusOK, temp)
}
