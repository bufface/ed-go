// Comentario de Línea
// Por convención en go, se utiliza para documentar el código

/*
Comentario multilinea
también conocidos como "de bloque"

Por convención, solo se utiliza para comentar código
*/

package main

import (
	"fmt"
)

func main() {
	// const nombre = "Cristian"

	/*
		  constantes no pormiten reasignación
			nombre = "Emanuel" -> No se permite
	*/

	// fmt.Println(nombre)

	// Tipos de datos (casteo)
	/* var a int
	var b int8
	n := "Cristian"
	p := "Buffa"

	a = 121212
	b = 5

	c := a + int(b)

	fmt.Println(c)
	fmt.Printf("Hola %s %s %d , ¿Cómo estás?\n", n, p, b)
	fmt.Printf("La variable b es de tipo: %T", b)

	// prioridad artimética
	// () % * / + -
	// d := 6 + 6*(6-6)
	// fmt.Println(d)
	*/

	// Valor 0 de las variables
	var nombre string  // valor 0 = ""
	var numero int     // valor 0 = 0
	var entiendes bool // valor 0 = false
	fmt.Println(nombre, numero, entiendes)
}
