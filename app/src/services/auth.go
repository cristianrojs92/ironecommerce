package services

import (
	"fmt"
	"../utils/logger"
	"context"
	"../models"
)

func InsertAuth(auth models.Auth) bool {

	function := "InsertAuth"
	var pg_conection PG_CONECTION

	//Nos conectamos a la base
	err := pg_conection.Connect()

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos")
		logger.Log(FILE, function, text)
		return false
	} 

	//Al finalizar la ejecucion se cierra la coneccion
	defer pg_conection.Close()

	//Formamos la consulta a la base
	query :=fmt.Sprintf("insert into auth (user_id,token) values (%d,'%s');", auth.UserId, auth.Token) 
	text := fmt.Sprintf("query = %s", query)
	logger.Log(FILE, function, text)

	//Realizamos la query
	row, err := pg_conection.conection_pool.Query(context.Background(), query)
	
	if err != nil {
		text := fmt.Sprintf("Error al ejectuar la query err = %s", err.Error())
		logger.Log(FILE, function, text)
		return false
	} 

	row.Close()

	return true
}

func DeleteAuth(user_id int) bool {

	function := "DeleteAuth"
	var pg_conection PG_CONECTION

	//Nos conectamos a la base
	err := pg_conection.Connect()

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos")
		logger.Log(FILE, function, text)
		return false
	} 

	//Al finalizar la ejecucion se cierra la coneccion
	defer pg_conection.Close()

	//Formamos la consulta a la base
	query :=fmt.Sprintf("delete from auth where user_id = %d;", user_id) 
	text := fmt.Sprintf("query = %s", query)
	logger.Log(FILE, function, text)

	//Realizamos la query
	row, err := pg_conection.conection_pool.Query(context.Background(), query)

	if err != nil {
		text := fmt.Sprintf("Error al ejectuar la query err = %s", err.Error())
		logger.Log(FILE, function, text)
		return false
	} 

	row.Close()

	return true
}