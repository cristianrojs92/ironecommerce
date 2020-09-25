package routers

//Dependencias
import (
	"github.com/gorilla/mux"
	"../controllers"
)

// Crea las rutas el modulo de usuarios
func setCategoriesRouters(router *mux.Router) *mux.Router {	
	router.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	return router
}
