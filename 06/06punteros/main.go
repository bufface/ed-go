package main

import "fmt"

func main() {
	a := 3
	fmt.Println("Antes de duplicar, a vale: ", a)
	duplicar(&a)
	fmt.Println("Después de duplicar, a vale: ", a)
}

func duplicar(a *int) {
	*a = *a * 2
	fmt.Println("Dentro de la función duplicar, a vale: ", *a)
}
