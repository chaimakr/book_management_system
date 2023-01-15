package handlers

import (
	"net/http"

	"github.com/chaimakr/book_management_system/core/getter/database"

	"github.com/gin-gonic/gin"
)

func GetBook(db database.BookInterface) gin.HandlerFunc {	
	return func(c *gin.Context) {
		id := c.Param("id")

		res, err := db.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
