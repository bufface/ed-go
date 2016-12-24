package main

import "fmt"

func main() {
	/*
		var nombre, apellido string

		nombre = "Cristian"
		apellido = "Buffa"
	*/

	// Asiganción miltiple
	nombre, apellido := "Cristian", "Buffa"

	// Reasignación de valores
	nombre = "José"

	/*
	  Solo puedo usar := para variables nuevas.
	  En éste caso, "nombre"" no es una variable nueva,
	  pero si lo es "edad"
	*/
	nombre, edad := "Emanuel", 29

	fmt.Println(nombre, apellido, edad)
}
