package main

import (
	"fmt"
)

// Persona es un tipo de estructura (equivalente a clases en otros lenguajes)
type Persona struct {
	Nombre string
	Edad   uint8
}

// Persona1 se agrega in slice a la estructura
type Persona1 struct {
	Nombre string
	Edad   uint8
	Emails []string
}

func main() {
	// Standart
	/*
		var persona1 Persona
		persona1.Nombre = "Cristian"
		persona1.Edad = 20

		fmt.Println(persona1.Nombre)
	*/

	// Shortcut
	/*
		persona2 := Persona{}
		persona2.Nombre = "Pablo"
		persona2.Edad = 30

		fmt.Println(persona2)
	*/

	/*
		persona3 := Persona{
			Nombre: "José",
			Edad:   66,
		}

		fmt.Println(persona3)
	*/

	/*
		persona4 := Persona{
			"Pedro",
			21,
		}

		fmt.Println(persona4)
	*/

	emails := []string{
		"bufface@gmail.com",
		"bufface@outlook.com",
	}

	persona5 := Persona1{
		"Cristian",
		29,
		emails,
	}

	fmt.Println(persona5)
}
