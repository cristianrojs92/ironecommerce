package routers

//Dependencias
import (
	"github.com/gorilla/mux"
	"../controllers"
)

// Crea las rutas el modulo de usuarios
func setUsersRouters(router *mux.Router) *mux.Router {	
	router.HandleFunc("/users", controllers.Autentication(controllers.GetUsers)).Methods("GET")
	return router
}
