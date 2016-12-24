// Comentario de Línea
// Por convención en go, se utiliza para documentar el código

/*
Comentario multilinea
también conocidos como "de bloque"

Por convención, solo se utiliza para comentar código
*/

package main

import "fmt"

func main() {
	const nombre = "Cristian"

	/*
		  constantes no pormiten reasignación
			nombre = "Emanuel" -> No se permite
	*/

	fmt.Println(nombre)
}
