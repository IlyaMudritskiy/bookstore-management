package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ilyamudritskiy/bookstore_management/pkg/routes"
)

func main() {
	var router = mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	fmt.Println("Server started on localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}