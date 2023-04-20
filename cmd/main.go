package main

import (
	"log"
	"net/http"

	"github.com/Harsh4902/Book-management-system/pkg/routes"
	"github.com/Harsh4902/Book-management-system/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8080",r))
}