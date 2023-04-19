package routes

import (
	"github.com/gorilla/mux"
	"github.com/Harsh4902/Book-management-system/pkg/controller"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.Handle("/book/",controller.CreateBook).Methods("POST")
	router.Handle("/book/",controller.GetBook).Methods("GET")
	router.Handle("/book/{bookId}",controller.GerBookById).Methods("GET")
	router.Handle("/book/{bookId}",controller.UpdateBook).Methods("PUT")
	router.Handle("/book/{bookId}",controller.DeleteBook).Methods("POST")
}
