package routes

import (
	"go-web/controllers"
	"net/http"
)

func LoadingRoutes() {
	// HandleFunc() = func http que recebe uma url e uma função à ser executada quando batemos nessa url
	// No caso, é chamada a func index ao batermos na url "/"
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new-product", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
