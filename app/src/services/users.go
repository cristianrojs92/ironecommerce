package services


import (
	"fmt"
	"../utils/logger"
	"context"
	"../models"
)


//Obtiene todos los usuario
func GetUsers() ([]models.User,error){

	function := "GetUsers"

	var pg_conection PG_CONECTION
	users := make([]models.User,0)

	//Nos conectamos a la base
	err := pg_conection.Connect()

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos")
		logger.Log("users.go", function, text)
		return nil,err
	} 

	//Al finalizar la ejecucion se cierra la coneccion
	defer pg_conection.Close()

	//Formamos la consulta a la base
	query := "select id,username,name,surname,email,password from users;"
	text := fmt.Sprintf("query = %s", query)
	logger.Log("users.go", function, text)

	//Realizamos la query
	rows, err := pg_conection.conection_pool.Query(context.Background(), query)

	//Si obtuvimos un error
	if err != nil {

		text := fmt.Sprintf("Error el obtener los usuarios")
		logger.Log("users.go", function, text)
		return nil,err
	}

	//Se recorren los resultados
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Surname, &user.Email, &user.Password)
		if err != nil {
			text := fmt.Sprintf("Error al obtener los usuarios err = %s", err.Error())
			logger.Log("users.go", function, text)
			break
		}
		
		users=append(users, user)
	}

	return users,nil

}

//Obtiene todos los usuario
func GetUser(user_name string) (models.User,error){

	var pg_conection PG_CONECTION
	var user models.User
	var valid_err error
	function := "GetUser"

	//Validamos datos de entrada
	if(user_name == "") {
		text := fmt.Sprintf("Nombre de usuario invalido")
		logger.Log(FILE, function, text)
		valid_err = fmt.Errorf("Nombre de usuario invalido")

		return user, valid_err
	}

	//Nos conectamos a la base
	err := pg_conection.Connect()

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos")
		logger.Log(FILE, function, text)
		return user,err
	} 

	//Al finalizar la ejecucion se cierra la coneccion
	defer pg_conection.Close()

	//Formamos la consulta a la base
	query := fmt.Sprintf("select id,username,name,surname,email,password from users where username='%s';", user_name)
	text := fmt.Sprintf("query = %s", query)
	logger.Log(FILE, function, text)

	//Realizamos la query
	err = pg_conection.conection_pool.QueryRow(context.Background(), query).Scan(&user.Id, &user.Username, &user.Name, &user.Surname, &user.Email, &user.Password)

	//Si obtuvimos un error
	if err != nil {

		text := fmt.Sprintf("Error el obtener los usuarios err = %s",err.Error())
		logger.Log(FILE, function, text)
		return user,err
	}

	return user ,nil
}