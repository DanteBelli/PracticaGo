package saludo

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hola(nombre string) (string, error) {
	if nombre == "" {
		return nombre, errors.New("Ingrese un Nombre")
	}
	message := fmt.Sprintf(randomFormat(), nombre)
	return message, nil
}
func randomFormat() string {
	formats := []string{
		"Hola %v, Bienvendio!",
		"Que bueno volver a ver %v,",
		"Vamos Boca %v,",
	}
	return formats[rand.Intn(len(formats))]
}
func Hellos(nombres []string) (map[string]string, error) {
	mensajes := make(map[string]string)
	for _, nombre := range nombres {
		mensaje, err := Hola(nombre)
		if err != nil {
			return nil, err
		}
		mensajes[nombre] = mensaje
	}
	return mensajes, nil
}
