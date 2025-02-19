package main

import (
	"fmt"
	"log"
	"practica/saludo"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := saludo.Hola("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
