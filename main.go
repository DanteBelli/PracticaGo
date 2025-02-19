package main

import (
	"fmt"
	"practica/saludo"
)

func main() {
	message := saludo.Hola("Dante")
	fmt.Println(message)
}
