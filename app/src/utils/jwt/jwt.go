package jwt

import (
 "crypto/hmac"
 "crypto/sha256"
 "encoding/base64"
 "encoding/json"
 "strings"
 "fmt"
 "errors"
 "time"
 "../../models"
)

//Se la estructura de payload
type Payload struct {
Username string
Password string
Exp int64
}

/*func main() {

   secret := "culo"


   t := time.Date(2020, time.November, 10, 23, 0, 0, 0, time.UTC)
   fmt.Printf("Go launched at %s\n", t.Local())
   fmt.Println("Unix", t.Unix())
   payload := Payload{
	Username: "cristian",
	Password: "1234",
	Exp: t.Unix(),
   }
   jwt:= Encode(payload,secret)
   fmt.Println("jwt ", jwt)
  
  res_pay,err :=  Decode(jwt, secret)

res_pay_2 := res_pay.(Payload)

  if err != nil {
   fmt.Println("Invalid payload:", err.Error())
  }	


  str_res, _ := json.Marshal(res_pay_2)
	
  fmt.Println("str_res ", str_res)
	
}*/



//Se recibe un string y se devuelve un string encodeado a base64
func Base64Encode(src string) string {
    return strings.
        TrimRight(base64.URLEncoding.
            EncodeToString([]byte(src)), "=")
}

// Base64Encode recibe un string encodeado a base 64 y lo devuelve desencodeado
func Base64Decode(src string) (string, error) {
    if l := len(src) % 4; l > 0 {
         src += strings.Repeat("=", 4-l)
    }
    decoded, err := base64.URLEncoding.DecodeString(src)
    if err != nil {
        errMsg := fmt.Errorf("Error al desencodear %s", err)
        return "", errMsg
    }
    return string(decoded), nil
}

//Retorna un hash generado en Hmac256 de tipo string usando palabra secreta
func Hash(src string, secret string) string {
  key := []byte(secret)
  h := hmac.New(sha256.New, key)
  h.Write([]byte(src))
  return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

//Comprueba si una firma es valida
func compareHmac(value string, hash string, secret string) bool {
 return hash == Hash(value, secret)
}

//Genera un token de tipo JWT
func Encode(payload models.PayloadUser, secret string) string {

 //Estructura para el header
 type Header struct {
 	Alg string `json:"alg"`
 	Typ string `json:"typ"`
 } 
 //Creamos el header
 header := Header{
 	Alg: "HS256",
 	Typ: "JWT",
 }

 //Se parsea el header a json
 json_header, _ := json.Marshal(header)

 //Encodeamos el header a base64
 base64_header := Base64Encode(string(json_header))
 
 //Parsea el payload a json
 json_payload, _ := json.Marshal(payload)

 //Encodeamos el payload a base64
 base64_payload := Base64Encode(string(json_payload))

 //Ser forma la firma con el header + el payload
 signatureValue := base64_header + "." + base64_payload

 //Se forma el token con la firma + la firma hasheada
 token := signatureValue + "." + Hash(signatureValue, secret)

 return token
}

//Desifra un token JWT
func Decode(jwt string, secret string) (bool, error) {

 //Se sepera el string : firma + hash de la firma
 token := strings.Split(jwt, ".")

 var payload models.PayloadUser

 //El string contiene  header +  payload + token
 if len(token) != 3 {
   split_err := errors.New("Token invalido: el mismo debe tener header, payload y palabra secreta")
   return false, split_err
 }

 //Desencodeamos el payload
 decoded_payload, payload_err := Base64Decode(token[1])
   if payload_err != nil {
   return false, fmt.Errorf("Payload invalido: %s", payload_err.Error())
 }

 //Almacenamos el payload
 parse_err := json.Unmarshal([]byte(decoded_payload), &payload)
 if parse_err != nil {
   return false, fmt.Errorf("Payload invalido: %s", parse_err.Error())
 }
 //Verificamos le token expiro
 if payload.Exp != 0 && time.Now().Unix() > payload.Exp {
   return false, errors.New("El token ha expirado")
 }

 //Formamos la firma
	signature_value := token[0] + "." + token[1]
	
 // Verificamos si la firma es valida
 if compareHmac(signature_value, token[2], secret) == false {
   return false, errors.New("Token invalido")
 }
 return true, nil
}
