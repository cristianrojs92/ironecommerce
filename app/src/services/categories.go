package services


import (
	"fmt"
	"../utils/logger"
	"context"
	"../models"
)

//Obtiene todas las categorias
func GetMainCategories() ([]models.Category,error){

	function := "GetMainCategories"

	var pg_conection PG_CONECTION
	categories := make([]models.Category,0)

	//Nos conectamos a la base
	err := pg_conection.Connect()

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos")
		logger.Log("categories.go", function, text)
		return nil,err
	} 

	//Al finalizar la ejecucion se cierra la coneccion
	defer pg_conection.Close()

	//Formamos la consulta a la base
	query := "SELECT id, main, name, icon_name FROM categories WHERE main=true;"
	text := fmt.Sprintf("query = %s", query)
	logger.Log("categories.go", function, text)

	//Realizamos la query
	rows, err := pg_conection.conection_pool.Query(context.Background(), query)

	//Si obtuvimos un error
	if err != nil {

		text := fmt.Sprintf("Error el obtener las categorias")
		logger.Log("categories.go", function, text)
		return nil,err
	}

	//Se recorren los resultados
	defer rows.Close()
	for rows.Next() {

		var category models.Category
		err := rows.Scan(&category.Id, &category.Main, &category.Name, &category.Icon_name)

		if err != nil {
			text := fmt.Sprintf("Error al obtener las categorias err = %s", err.Error())
			logger.Log("categories.go", function, text)
			break
		}
		
		categories=append(categories, category)
	}

	return categories,nil

}

//Obtiene todas las categorias relacionadas a una principal
func GetSubCategories(id_main_category int) ([]models.Category,error){

	function := "GetSubCategories"

	var pg_conection PG_CONECTION
	categories := make([]models.Category,0)

	//Nos conectamos a la base
	err := pg_conection.Connect()

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos")
		logger.Log("categories.go", function, text)
		return nil,err
	} 

	//Al finalizar la ejecucion se cierra la coneccion
	defer pg_conection.Close()

	//Formamos la consulta a la base
	
	query := fmt.Sprintf("SELECT c.id, c.main, c.name, c.icon_name from categories_relations AS cr INNER JOIN categories as c ON  cr.id_child = c.id where cr.id_main = %d;",id_main_category)
	text := fmt.Sprintf("query = %s", query)
	logger.Log("categories.go", function, text)

	//Realizamos la query
	rows, err := pg_conection.conection_pool.Query(context.Background(), query)

	//Si obtuvimos un error
	if err != nil {

		text := fmt.Sprintf("Error el obtener las sub categorias")
		logger.Log("categories.go", function, text)
		return nil,err
	}

	//Se recorren los resultados
	defer rows.Close()
	for rows.Next() {

		var category models.Category
		err := rows.Scan(&category.Id, &category.Main, &category.Name, &category.Icon_name)

		if err != nil {
			text := fmt.Sprintf("Error al obtener las categorias err = %s", err.Error())
			logger.Log("categories.go", function, text)
			break
		}
		
		categories=append(categories, category)
	}

	return categories,nil

}

