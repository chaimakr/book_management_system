package handlers

import (
	"net/http"

	"github.com/chaimakr/book_management_system/core/setter/database"
	"github.com/chaimakr/book_management_system/core/setter/models"

	"github.com/gin-gonic/gin"
)

func InsertBook(db database.BookInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		book := models.Book{}
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Insert(book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
