package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"	
	"../utils/logger"
	"../utils/httpResponse"
	"../services"
)

//Obtiene todas las categorias
func GetCategories(w http.ResponseWriter, r *http.Request){

	function := "GetCategories";

	//Obtenemos todas las categorias principales
	categories,err := services.GetMainCategories();

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Obtuvimos un error al obtener las categorias")
		logger.Log("categories.go", function, text)
		httpResponse.Error(w, "Ocurrio un error", 500)
		
	} else {
		text := fmt.Sprintf("Se obtienen las categorias correctamente")
		logger.Log("categories.go", function,text)

	//Verificamos si tenemos categorias
	if(categories != nil && len(categories) > 0) {

		//Obtenemos las subcategorias de las categorias principales
		for index, category := range categories {
			
			//Obtenemos todas las subcategorias
			sub_categories,_ := services.GetSubCategories(category.Id);

			//Verificamos si tenemos subcatgorias
			if(sub_categories != nil && len(sub_categories) > 0) {
				categories[index].Sub_categories = sub_categories;
			}
		}
	}

		//Formamos la respuesta
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(categories)
		w.WriteHeader(http.StatusOK)
		w.Write(j)

	}
}