package main

import "fmt"

func main() {
	/*
		anonima := func() {
			fmt.Println("Me imprimo desde una variable")
		}

		anonima()
	*/

	secuencia1 := secuencia()

	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())

	fmt.Println("-------------------------")

	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
}

// "secuencia" va a retornar una funcion anonima
// y ésta funcion anonima va a retornar un "int32""
func secuencia() func() int32 {
	var x int32

	return func() int32 {
		x++
		return x
	}
}
