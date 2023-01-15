package main

import (
	"context"
	"fmt"

	"github.com/chaimakr/book_management_system/core/getter/config"
	"github.com/chaimakr/book_management_system/core/getter/database"
	"github.com/chaimakr/book_management_system/core/getter/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()
	fmt.Println(conf)
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	fmt.Println(conf, collection, db)

	client := &database.BookClient{
		Col: collection,
		Ctx: ctx,
	}

	r := gin.Default()

	todos := r.Group("/books")
	{
		todos.GET("/", handlers.SearchBooks(client))
		todos.GET("/:id", handlers.GetBook(client))
		todos.POST("/add", handlers.InsertBook(client))
		todos.PATCH("/:id", handlers.UpdateBook(client))
		todos.DELETE("/:id", handlers.DeleteBook(client))
	}

	r.Run(":8080")
}
