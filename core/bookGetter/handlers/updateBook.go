package controllers

import (
	"net/http"

	"github.com/chaimakr/book_management_system/core/getter/database"

	"github.com/gin-gonic/gin"
)

func UpdateBook(db database.TodoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book interface{}
		id := c.Param("id")
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Update(id, book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
