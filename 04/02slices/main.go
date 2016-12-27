package main

import (
	"fmt"
)

func main() {
	// Slices
	// 1. Apuntador a un array
	// 2. Tienen un tamaño (no fijo)
	// 3. Tienen una capacidad
	// var nombres []string

	// Shortcut
	// make(tipo de dato, tamaño inicial [,capacidad inicial])
	nombres := []string{
		"Cristian",
		"Emanuel",
		"Buffa",
	}
	/*
		nombres := make([]string, 0)
		fmt.Printf("Tamaño: %d. Capacidad: %d\n", len(nombres), cap(nombres))

		nombres = append(nombres, "Cristian")
		fmt.Printf("Tamaño: %d. Capacidad: %d\n", len(nombres), cap(nombres))

		nombres = append(nombres, "Emanuel")
		fmt.Printf("Tamaño: %d. Capacidad: %d\n", len(nombres), cap(nombres))

		nombres = append(nombres, "Buffa")
		fmt.Printf("Tamaño: %d. Capacidad: %d\n", len(nombres), cap(nombres))

		nombres = append(nombres, "Franco")
		fmt.Printf("Tamaño: %d. Capacidad: %d\n", len(nombres), cap(nombres))

		nombres = append(nombres, "Joaquin")
		fmt.Printf("Tamaño: %d. Capacidad: %d\n", len(nombres), cap(nombres))

		nombres[3] = "Francisco"
	*/
	fmt.Println(nombres)
}
