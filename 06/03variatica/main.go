package main

import "fmt"

func main() {
	saludarVarios(18, "Cristian", "Manolo", "Sol", "Jimena", "Pedrito")
}

func saludarVarios(edad uint8, name ...string) {
	fmt.Printf("%T\n", name)

	for _, val := range name {
		fmt.Println("Hola", val, "tienes", edad)
	}
}
