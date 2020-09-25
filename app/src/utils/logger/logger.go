package logger

import (
		"log"
		"fmt"
)

//Logea el texto pasado por parametro
func Log(file string, function string, text string) {

   value := fmt.Sprintf("[%s] [%s]: %s", file, function, text)
		log.Print(value);
}

//Logea el error pasado por parametro
func Error(file string, function string, text string) {
	
	value := fmt.Sprintf("[%s] [%s]: %s", file, function, text)
	log.Fatalf(value);
}
