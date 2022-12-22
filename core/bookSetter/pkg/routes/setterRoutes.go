package routes

import (
	"github.com/chaimakr/book_management_system/bookGetter/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookSetterRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
