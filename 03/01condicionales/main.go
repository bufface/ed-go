package main

import (
	"fmt"
)

// \\ Or
// && And
// ! Not
func main() {
	// isValid := true

	if edad, nombre := 15, "Cristian"; edad < 14 {
		fmt.Println(nombre, "Es un niño")
	} else if edad < 18 {
		fmt.Println(nombre, "Es un adolescente")
	} else if edad < 30 {
		fmt.Println("Es un adulto")
	} else {
		fmt.Println("Es un adulto mayor")
	}

	/*if isValid {
		fmt.Println("Condicion verdadera")
		fmt.Println(isValid)

		if 5 < 10 {
			fmt.Println("5 es menor a 10")
		} else {
			fmt.Println("5 no es mayor a 10")
		}
	} else {
		fmt.Println("Condicion falsa")
		fmt.Println(isValid)
	}*/

	fmt.Println("Fin del programa")
}
