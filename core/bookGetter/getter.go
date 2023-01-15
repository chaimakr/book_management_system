package main

import (
	"context"

	"github.com/chaimakr/book_management_system/core/getter/config"
	"github.com/chaimakr/book_management_system/core/getter/database"
	"github.com/chaimakr/book_management_system/core/getter/handlers"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func main() {

	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)
	client := &database.BookClient{
		Col: collection,
		Ctx: ctx,
	}

	r := gin.Default()

	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// set middleware for gin
	m.Use(r)
	todos := r.Group("/books")
	{
		todos.GET("/", handlers.SearchBooks(client))
		todos.GET("/:id", handlers.GetBook(client))
	}

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	r.Run(":8081")
}
