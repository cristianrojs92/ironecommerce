package config

//Paquetes
import (
	"io"
	"os"
	"encoding/json"
  "../utils/logger"
)

var FILE = "config.go"

type CONFIG struct {

	//Servidor
	Server SERVER
	Db DB

}

type SERVER struct {

	// Puerto
	Port int
	Secret string
}

type DB struct {
	Name string
	User string
	Password string
	Host string
}

//Variables
var config CONFIG

// Inicializa la configuracion
func InitConfig() {

	function := "InitConfig";

	//Realizamos la lectura de la configuracion.
	file, err := os.Open("dist/etc/config.json");

	//Al finalizar la ejecucion de esta funcion cerramos la lectura del archivo
	defer file.Close();
	
 //Si obtuvimos un error
 if err != nil {
	 logger.Error(FILE, function, err.Error())
 }
	
 if file == nil {
	
  err := "No se pudo realizar la lectura del archivo";
	logger.Error(FILE, function, err);
 }
	
  //Deciframos el json completo.
	decoder := json.NewDecoder(file);
		                
	for {
		if err := decoder.Decode(&config); err == io.EOF {
			break
		} else if err != nil {
			logger.Error(FILE, function, err.Error());
		}
	}              
	
}

/**
	Retorna la configuracion actual
**/
func GetConfig() (CONFIG){

	return config;

} 
