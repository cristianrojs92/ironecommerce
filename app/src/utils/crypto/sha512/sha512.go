package sha512

import (	
	"fmt"
	"crypto/sha512"
)

//Logeo de un usuario
func GenerateHash(text string) string{

	hash := ""

	//Validamos datos de entrada
	if(text != "") {

		//Generamos un hasheador sha512
		hasher := sha512.New()

		//Convertimos el valos entranta en un array de bytes
		byte_text:= []byte(text)

		//Agregamos los bytes en el hasheador
		hasher.Write(byte_text)

		//Obtenemos el hash
		hash = fmt.Sprintf("%x",hasher.Sum(nil))
	}
  
	return hash	
}