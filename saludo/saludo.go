package saludo

import (
	"fmt"
)

func Hola(nombre string) string {
	message := fmt.Sprintf("Hi , %v. Welcome!", nombre)
	return message
}
