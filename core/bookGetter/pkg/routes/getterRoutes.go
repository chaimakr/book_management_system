package routes

import (
	"github.com/chaimakr/book_management_system/bookGetter/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookGetterRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")

}
