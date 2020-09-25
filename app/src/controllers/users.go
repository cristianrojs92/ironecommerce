package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"	
	"../utils/logger"
	"../utils/httpResponse"
	"../services"
)

//Obtiene todos los usuario
func GetUsers(w http.ResponseWriter, r *http.Request){

	function := "GetUsers";

	//Obtenemos todos los usuarios
	users,err := services.GetUsers();

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Obtuvimos un error al obtener el usuario")
		logger.Log("users.go", function, text)
		httpResponse.Error(w, "Ocurrio un error", 500)
		
	} else {
		text := fmt.Sprintf("Se obtienen los usuarios correctamente")
		logger.Log("users.go", function,text)

	
		//Formamos la respuesta
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(users)
		w.WriteHeader(http.StatusOK)
		w.Write(j)

	}
}