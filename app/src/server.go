package main

import (
  "net/http"
  "fmt"
  "os"
  "../src/utils/logger"
  "../src/config"
  "../src/routes"
)

var FILE = "server.go"

func main() {

  function := "main";

  //Caragamos las configuraciones del proyecto.
  config.InitConfig()
  configuration := config.GetConfig()
  
  //Inicializamos las rutas del servidor
  router := routers.InitRoutes()

  //Direccion y puerto donde correra el servidor
  addr := fmt.Sprintf("127.0.0.1:%d", configuration.Server.Port)

  //Creamos el servidor
	server := &http.Server{
    Addr:    addr,
		Handler: router,
  }

  //Pid del proceso
  text := fmt.Sprintf("Proceso corriendo, PID =  %d", os.Getpid())
  logger.Log(FILE, function,text)
  
  //Incializamos le servidor
  text = fmt.Sprintf("Se inicia el servidor en %s", addr)
  logger.Log(FILE, function,text)
  err := server.ListenAndServe()

  //Si obtuvimos un error.
  if(err != nil) {
    text := fmt.Sprintf("Error al iniciar el servidor en %s", err.Error())
    logger.Error(FILE, function,text)
  }
}