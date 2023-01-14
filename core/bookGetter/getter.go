package main

import (
	"fmt"
	"net/http"

	"github.com/chaimakr/book_management_system/core/getter/config"
	"github.com/chaimakr/book_management_system/core/getter/database"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)
	db := database.ConnectDB(conf.Mongo)
	fmt.Println(db)
	r := mux.NewRouter()
	http.ListenAndServe(":8080", r)
}
