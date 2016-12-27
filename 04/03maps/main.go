package main

import (
	"fmt"
)

func main() {
	// Maps (algo así como arrays asociativos)
	// idiomas["es"] = "Español"

	/*
		idiomas := make(map[string]string)

		idiomas["es"] = "Español"
		idiomas["en"] = "Inglés"
		idiomas["fr"] = "Francés"
	*/
	/*
		idiomas := map[string]string{
			"es": "Español",
			"en": "Inglés",
			"fr": "Francés",
			"pt": "Portugués",
		}
	*/
	/*
		fmt.Println(idiomas)
		delete(idiomas, "en")
		fmt.Println(idiomas)
	*/

	/*
		if idioma, ok := idiomas["pt"]; ok {
			fmt.Println("Existe la posición:", idioma)
		} else {
			fmt.Println("No Existe la posición")
		}
	*/

	numeros := map[int]string{
		1:    "Número 1",
		2016: "Número 2",
		-6:   "Número 3",
	}

	fmt.Println(numeros)
}
