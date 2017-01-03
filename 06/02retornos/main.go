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

	/*
		var edad uint8
		edad = 30

		fmt.Println(tipoEdad(edad))
	*/

	n := []int8{52, 12, 78, 94, 83, 19, -77}
	maximo, minimo := maxymin(n)

	fmt.Println("Máximo:", maximo)
	fmt.Println("Mínimo:", minimo)
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

func maxymin(numeros []int8) (max int8, min int8) {

	for _, value := range numeros {

		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return
}
