package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	
	//Seteamos las rutas
	router = setUsersRouters(router)
	router = setAuthRouters(router)
	router = setCategoriesRouters(router)
	router = setProductsRouters(router)

	return router
}