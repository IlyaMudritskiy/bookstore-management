package routes

import (
	"github.com/gorilla/mux"
	"github.com/ilyamudritskiy/bookstore_management/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	
}

func main() {

}