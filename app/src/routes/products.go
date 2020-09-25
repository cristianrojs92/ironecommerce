package routers

//Dependencias
import (
	"github.com/gorilla/mux"
	"../controllers"
)

// Crea las rutas el modulo de usuarios
func setProductsRouters(router *mux.Router) *mux.Router {	
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products/offers", controllers.GetOffers).Methods("GET")
	return router
}
