package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"	
	"../utils/logger"
	"../utils/httpResponse"
	"../services"
)

//Obtiene todos los productos
func GetProducts(w http.ResponseWriter, r *http.Request){

	function := "GetProducts";

	//Obtenemos todos los productos
	products,err := services.GetProducts();

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Obtuvimos un error al obtener los productos")
		logger.Log("products.go", function, text)
		httpResponse.Error(w, "Ocurrio un error", 500)
		
	} else {
		text := fmt.Sprintf("Se obtienen los productos correctamente")
		logger.Log("products.go", function,text)

	
		//Formamos la respuesta
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(products)
		w.WriteHeader(http.StatusOK)
		w.Write(j)

	}
}

//Obtiene todas las ofertas
func GetOffers(w http.ResponseWriter, r *http.Request){

	function := "GetOffers";

	//Obtenemos todos los productos
	offers,err := services.GetOffers();

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Obtuvimos un error al obtener las ofertas")
		logger.Log("products.go", function, text)
		httpResponse.Error(w, "Ocurrio un error", 500)
		
	} else {
		text := fmt.Sprintf("Se obtienen las ofertas correctamente")
		logger.Log("products.go", function,text)

	
		//Formamos la respuesta
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(offers)
		w.WriteHeader(http.StatusOK)
		w.Write(j)

	}
}