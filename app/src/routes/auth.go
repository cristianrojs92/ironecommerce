package routers

//Dependencias
import (
	"github.com/gorilla/mux"
	"../controllers"
)

// Crea las rutas el modulo de usuarios
func setAuthRouters(router *mux.Router) *mux.Router {	
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	return router
}