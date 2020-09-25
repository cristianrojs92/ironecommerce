package controllers

import (
	"encoding/json"
	"net/http"	
	"fmt"
	"time"
	"../models"
	"../services"
	"../utils/logger"
	"../utils/httpResponse"
	"../utils/crypto/sha512"
	"../utils/jwt"
	"../config"
)

//Logeo de un usuario
func Login(w http.ResponseWriter, r *http.Request){

	function := "Login";

	//Parseamos el body a json
	decoder := json.NewDecoder(r.Body)
	var login models.LoginRequest 
	err := decoder.Decode(&login)
	if err != nil {
		text := fmt.Sprintf("Body invalido")
		logger.Log("auth.go", function, text)
		httpResponse.Error(w, "Body invalido", 500)
	} else {

		//Verificamos que los campos no estan vacios
		if login.Username == "" && login.Password == "" {
			text := fmt.Sprintf("Datos faltantes")
			logger.Log("auth.go", function, text)
			httpResponse.Error(w, "Datos faltantes", 500)

		} else {

			//Verificamos si el usuario existe
			user,err:= services.GetUser(login.Username)
			if err != nil {
				text := fmt.Sprintf("Error al obtener el usuario")
				logger.Log("auth.go", function, text)
				httpResponse.Error(w, "El usuario no existe", 401)
			}

			//Generamos el hash de la contraseña
			hash_password := sha512.GenerateHash(login.Password)
			
			//Comparamos las contraseñas
			if(hash_password != user.Password) {
				text := fmt.Sprintf("Contraseña invalida")
				logger.Log("auth.go", function, text)
				httpResponse.Error(w, "Contraseña invalida", 401)
			} else {

				//El token expira en 7 dias
				now := time.Now()

				exp := now.AddDate(0,0,7).Unix()

				payload := models.PayloadUser {Id : user.Id, Username : user.Username, Name : user.Name, Surname : user.Surname, Email : user.Email, Exp : exp }
				
				//Obtenemos la palabra secreta
				config := config.GetConfig();
				secret := config.Server.Secret

				//Generamos el token con palabra secreta
				token :=  jwt.Encode(payload, secret)

				//Se elimina el token anterior
				resp := services.DeleteAuth(user.Id)

				if(resp != true) {
					text := fmt.Sprintf("Error al borrar la autenticacion anterior del usuario")
					logger.Log("auth.go", function, text)
					httpResponse.Error(w, "Error al borrar la autenticacion anterior del usuario", 500)
				} else {

					//Almanemos el token en la base
					var auth models.Auth = models.Auth {
						UserId : user.Id,
						Token : token,
					}
					resp := services.InsertAuth(auth)

					if(resp != true) {
						text := fmt.Sprintf("Error al insertar la autenticacion del usuario")
						logger.Log("auth.go", function, text)
						httpResponse.Error(w, "Error al insertar la autenticacion del usuario", 500)
					} else {

						//Generamos la respuesta del login
						login_response := models.LoginResponse{Token: token, Id : user.Id, Username : user.Username, Name : user.Name, Surname : user.Surname, Email : user.Email} 

						//Formamos la respuesta
						w.Header().Set("Content-Type", "application/json")
						j, _ := json.Marshal(login_response)
						w.WriteHeader(http.StatusOK)
						w.Write(j)

					}
				}
			}
		}
	}
}

//Se verifica la validez del token
func Autentication(next http.HandlerFunc) http.HandlerFunc {
	
  return func(w http.ResponseWriter, r *http.Request) {
		function := "Autentication";

		//Obtenemos el token
		token := r.Header.Get("Authorization")

		//Obtenemos la palabra secreta
		config := config.GetConfig();
		secret := config.Server.Secret

		//varificamos si el token es valido
		resp,err := jwt.Decode(token, secret)

		if err != nil || resp == false {

			text := fmt.Sprintf("Token invalido err = %s", err.Error())
			logger.Log("auth.go", function, text)
			httpResponse.Error(w, "Token invalido", 401)
		} else {
			//Ejecutamos el siguiente middleware
			next(w, r)
		}
  }
}

