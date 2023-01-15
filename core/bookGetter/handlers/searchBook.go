package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chaimakr/book_management_system/core/getter/database"
	utils "github.com/chaimakr/book_management_system/core/getter/utils"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func SearchBooks(db database.BookInterface, r *http.Request) gin.HandlerFunc {

	return func(c *gin.Context) {
		logger := utils.BuildLogger()

		tracer := utils.BuildTracer()

		rid := utils.GetRequestID(r)
		_, span := tracer.Start(r.Context(), "handle", trace.WithAttributes(
			attribute.String("request_id", rid), attribute.String("client_ip", r.RemoteAddr),
		))
		defer span.End()
		requestLogger := logger.With("client_ip", r.RemoteAddr, "request_id", rid)
		var filter interface{}
		query := c.Query("q")

		if query != "" {
			err := json.Unmarshal([]byte(query), &filter)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
				requestLogger.Errorw("Failed to retrieve books")
				return
			}
		}

		res, err := db.Search(filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			requestLogger.Errorw("Failed to retrieve books")
			return
		} else {
			requestLogger.Infow("books retrieved successfully")
		}

		c.JSON(http.StatusOK, res)
	}
}
