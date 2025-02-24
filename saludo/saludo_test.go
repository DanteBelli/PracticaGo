package saludo

import (
	"regexp"
	"testing"
)

func TestHolaNombre(t *testing.T) {
	nombre := "Dante"
	quiero := regexp.MustCompile(`\b` + nombre + `\b`)
	msg, err := Hola("Dante")
	if !quiero.MatchString(msg) || err != nil {
		t.Fatalf(`Hola("Dante") = %q , %v , no es igual con %#q , nil`, msg, err, quiero)
	}
}
func TestHolaVacio(t *testing.T) {
	msg, err := Hola("")
	if msg != "" || err == nil {
		t.Fatalf(`Hola("") = %q , %v , quiero "" , error`, msg, err)
	}
}
