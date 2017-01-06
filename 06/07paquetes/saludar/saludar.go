package saludar

import "fmt"

// MeVes ...
var MeVes string

var noMeVes string

// Saludar saluda a una persona
func Saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func noVisible() {
	fmt.Println(noMeVes)
}
