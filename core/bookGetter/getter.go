package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chaimakr/book_management_system/core/getter/config"
	"github.com/chaimakr/book_management_system/core/getter/database"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()
	fmt.Println(conf)
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	fmt.Println(conf, collection, db)

	// client := &database.TodoClient{
	// 	Col: collection,
	// 	Ctx: ctx,
	// }

	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)
}
