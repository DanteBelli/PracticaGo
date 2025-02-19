package saludo

import (
	"errors"
	"fmt"
)

func Hola(nombre string) (string, error) {
	if nombre == "" {
		return "", errors.New("Ingrese un Nombre")
	}
	message := fmt.Sprintf("Hi , %v. Welcome!", nombre)
	return message, nil
}
