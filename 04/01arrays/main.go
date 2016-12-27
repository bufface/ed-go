package main

import (
	"fmt"
)

func main() {
	// Arrays
	// Todos los valores son del mismo tipo
	// Tamaño fijo
	/*
		var nombres [3]string
		nombres[0] = "Cristian"
		nombres[1] = "Emanuel"
		nombres[2] = "Buffa"
		// nombres[3] = "Fuera de los limites"
	*/

	nombres := [3]string{
		"Cristian",
		"Emanuel",
		"Buffa",
	}
	nombre := nombres[1]
	size := len(nombres)

	fmt.Println("Tamaño del array nombres:", size)
	fmt.Println(nombre)
	// fmt.Println(nombres[inicio:final(excluyente)]) O sea, que no toma en cuenta el valor en "final"
	fmt.Println(nombres[0:2])
}
