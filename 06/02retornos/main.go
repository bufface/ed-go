package main

import "fmt"

func main() {
	/*
		var numero1, numero2 int8

		numero1 = 15
		numero2 = 3

		result := sumar(numero1, numero2)

		fmt.Println(result)
	*/

	var edad uint8
	edad = 30

	fmt.Println(tipoEdad(edad))
}

func sumar(a, b int8) int8 {
	return a + b
}

func tipoEdad(edad uint8) string {
	var tipo string

	switch {
	case edad < 12:
		tipo = "Niño"
	case edad < 18:
		tipo = "Adolescente"
	default:
		tipo = "Adulto"
	}

	return tipo
}
