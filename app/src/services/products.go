package services


import (
	"fmt"
	"../utils/logger"
	"context"
	"../models"
)

//Obtiene todas las categorias
func GetProducts() ([]models.Product,error){

	function := "GetProducts"

	var pg_conection PG_CONECTION
	products := make([]models.Product,0)

	//Nos conectamos a la base
	err := pg_conection.Connect()

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos")
		logger.Log("products.go", function, text)
		return nil,err
	} 

	//Al finalizar la ejecucion se cierra la coneccion
	defer pg_conection.Close()

	//Formamos la consulta a la base
	query := "select id, id_category,name,description,price,stock,image from products;"
	text := fmt.Sprintf("query = %s", query)
	logger.Log("products.go", function, text)

	//Realizamos la query
	rows, err := pg_conection.conection_pool.Query(context.Background(), query)

	//Si obtuvimos un error
	if err != nil {

		text := fmt.Sprintf("Error al obtener los productos")
		logger.Log("products.go", function, text)
		return nil,err
	}

	//Se recorren los resultados
	defer rows.Close()
	for rows.Next() {



		var product models.Product
		err := rows.Scan(&product.Id, &product.Id_category, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Image)

		if err != nil {
			text := fmt.Sprintf("Error al obtener los productos err = %s", err.Error())
			logger.Log("products.go", function, text)
			break
		}
		
		products=append(products, product)
	}

	return products,nil

}


//Obtiene todas las ofertas
func GetOffers() ([]models.Product,error){

	function := "GetOffers"

	var pg_conection PG_CONECTION
	products := make([]models.Product,0)

	//Nos conectamos a la base
	err := pg_conection.Connect()

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos")
		logger.Log("products.go", function, text)
		return nil,err
	} 

	//Al finalizar la ejecucion se cierra la coneccion
	defer pg_conection.Close()

	//Formamos la consulta a la base
	query := "select p.id, p.id_category,p.name,p.description,p.price,p.stock,p.image from products as p inner join offers as o on o.id_product = p.id ;"
	text := fmt.Sprintf("query = %s", query)
	logger.Log("products.go", function, text)

	//Realizamos la query
	rows, err := pg_conection.conection_pool.Query(context.Background(), query)

	//Si obtuvimos un error
	if err != nil {

		text := fmt.Sprintf("Error al obtener los productos")
		logger.Log("products.go", function, text)
		return nil,err
	}

	//Se recorren los resultados
	defer rows.Close()
	for rows.Next() {



		var product models.Product
		err := rows.Scan(&product.Id, &product.Id_category, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Image)

		if err != nil {
			text := fmt.Sprintf("Error al obtener los productos err = %s", err.Error())
			logger.Log("products.go", function, text)
			break
		}
		
		products=append(products, product)
	}

	return products,nil

}

