package main

import (
	"context"
	"net/http"

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
	req := http.Request{}
	r := gin.Default()

	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/query/metrics")

	// set middleware for gin
	m.Use(r)
	todos := r.Group("/query")
	{
		todos.GET("/", handlers.SearchBooks(client, &req))
		todos.GET("/:id", handlers.GetBook(client, &req))
	}

	r.Run(":8081")
}
