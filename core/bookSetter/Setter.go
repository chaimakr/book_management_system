package main

import (
	"context"

	"github.com/chaimakr/book_management_system/core/setter/config"
	"github.com/chaimakr/book_management_system/core/setter/database"
	"github.com/chaimakr/book_management_system/core/setter/handlers"
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

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/command/metrics")

	// set middleware for gin
	m.Use(r)
	todos := r.Group("/command")
	{
		todos.POST("/add", handlers.InsertBook(client))
		todos.PATCH("/:id", handlers.UpdateBook(client))
		todos.DELETE("/:id", handlers.DeleteBook(client))
	}

	r.Run(":8080")
}
