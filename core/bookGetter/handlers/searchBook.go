package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chaimakr/book_management_system/core/getter/database"
	"github.com/gin-gonic/gin"
)

func SearchBooks(db database.BookInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter interface{}
		query := c.Query("q")

		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				//g.GetMonitor().GetMetric("totalRequest").Inc([]string{"route_/books", "status_400", "requestType_GET"})
				return
			}
		}

		res, err := db.Search(filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			//g.GetMonitor().GetMetric("totalRequest").Inc([]string{"route_/books", "status_400", "requestType_GET"})
			return
		} else {
			//g.GetMonitor().GetMetric("totalRequest").Inc([]string{"route_/books", "status_200", "requestType_GET"})
		}

		c.JSON(http.StatusOK, res)
	}
}