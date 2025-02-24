package main

import (
	"fmt"
	"log"
	"practica/saludo"
)

func main() {
	log.SetPrefix("Saludos: ")
	log.SetFlags(0)
	nombre := []string{"Dante", "Rosario", "Pedro"}
	message, err := saludo.Hellos(nombre)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
