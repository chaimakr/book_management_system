package handlers

import (
	"net/http"

	"github.com/chaimakr/book_management_system/core/setter/database"
	"github.com/chaimakr/book_management_system/core/setter/models"
	utils "github.com/chaimakr/book_management_system/core/setter/utils"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func InsertBook(db database.BookInterface, r *http.Request) gin.HandlerFunc {
	return func(c *gin.Context) {
	logger := utils.BuildLogger()

	tracer := utils.BuildTracer()

	rid := utils.GetRequestID(r)
	_, span := tracer.Start(r.Context(), "handle", trace.WithAttributes(
		attribute.String("request_id", rid), attribute.String("client_ip", r.RemoteAddr),
	))
	defer span.End()
	requestLogger := logger.With("client_ip", r.RemoteAddr, "request_id", rid)
		book := models.Book{}
		err := c.BindJSON(&book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			requestLogger.Errorw("Failed to add the book")
			return
		}

		res, err := db.Insert(book)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
		requestLogger.Infow("book added successfully")

	}
}
