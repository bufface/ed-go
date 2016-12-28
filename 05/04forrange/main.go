package main

import "fmt"

func main() {
	// numeros := []string{"cero", "uno", "dos", "tres"}

	/*
		for posicion, valor := range numeros {
			fmt.Println(posicion, valor)
		}
	*/

	/*
		for _, valor := range numeros {
			fmt.Println(valor)
		}
	*/

	/*
		for indice := range numeros {
			fmt.Println(indice)
		}
	*/

	/*
			nombres := map[string]string{
				"es": "España",
				"br": "Brazil",
				"ar": "Argentina",
			}

		for posicion, valor := range nombres {
			fmt.Println(posicion, valor)
		}
	*/

	/*
		frase := "Hola mundo"

		for i, v := range frase {
			fmt.Println(i, string(v)) // El valor está casteado
		}
	*/

	for _, v := range []int{35, 56, 23} {
		fmt.Println(v)
	}
}
