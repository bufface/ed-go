package main

import "fmt"

func saludar(nombre string, edad uint8) {
	fmt.Printf("Hola %s, tienes %d años de edad\n", nombre, edad)
}

func main() {
	saludar("Cristian", 29)
}
