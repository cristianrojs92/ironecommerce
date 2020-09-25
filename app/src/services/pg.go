package services

import (
	"context"
	"fmt"	
	"../utils/logger"
	"../config"
	"github.com/jackc/pgx/pgxpool"
)

var FILE = "pg.go"

var init_config = false

var conection_string string

type PG_CONECTION struct {
	
	conection_pool *pgxpool.Pool
}


//Inicializa las configuraciones
func (c *PG_CONECTION) initConfig() {

	function := "initConfig";

	//Obtenemos la configuracion
	configuration := config.GetConfig()

	//Formamos el string de conexion
	if( configuration.Db.User == "" || configuration.Db.Host == "" || configuration.Db.Name == "" ) {
	 
		//Si hay datos nulos se utiliza un string de coneccion por defecto
		conection_string = "postgres://ironuser:@localhost/irondb"
		text := fmt.Sprintf("Configuracioens invalidas se utiliza string de conexion por defecto")
		logger.Log(FILE, function,text)

	} else {
		conection_string = fmt.Sprintf("postgres://%s:%s@%s/%s", configuration.Db.User, configuration.Db.Password, configuration.Db.Host, configuration.Db.Name)
	}

	init_config = true

}

//Se realiza la conexion con la base de datos
func (c *PG_CONECTION) Connect() (error) {

	function := "connect";
	var err error

	//Verificamos si se inicializaron las configuraciones de al db
	if (init_config == false) {
		c.initConfig();
	}

	//Obtenemos el contexto
	ctx := context.Background()

	//Realizammos el pool de conexion
	c.conection_pool, err = pgxpool.Connect(ctx, conection_string)

	//Si obtuvimos un error
	if err != nil {
		text := fmt.Sprintf("Error al conectarse a la base de datos. %s err = %s", conection_string, err.Error())
		logger.Log(FILE, function, text)
		return err
	}

	return nil
}

//Se cierra la coneccion
func (c *PG_CONECTION) Close() {
	c.conection_pool.Close()
}
