package handlers

import (
	"net/http"

	"github.com/chaimakr/book_management_system/core/getter/database"
	logger := utils.BuildLogger()

	tracer := utils.BuildTracer()

	rid := utils.GetRequestID(r)
	_, span := tracer.Start(r.Context(), "handle", trace.WithAttributes(
		attribute.String("request_id", rid), attribute.String("client_ip", r.RemoteAddr),
	))
	defer span.End()
	requestLogger := logger.With("client_ip", r.RemoteAddr, "request_id", rid)
)

func GetBook(db database.BookInterface, r *http.Request) gin.HandlerFunc {
	logger := utils.BuildLogger()

	tracer := utils.BuildTracer()

	rid := utils.GetRequestID(r)
	_, span := tracer.Start(r.Context(), "handle", trace.WithAttributes(
		attribute.String("request_id", rid), attribute.String("client_ip", r.RemoteAddr),
	))
	defer span.End()
	requestLogger := logger.With("client_ip", r.RemoteAddr, "request_id", rid)
	return func(c *gin.Context) {
		id := c.Param("id")

		res, err := db.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			requestLogger.Errorw("Failed to retrieve the book")
			return
		}

		c.JSON(http.StatusOK, res)
		requestLogger.Infow("book retrieved successfully")
	}
}
