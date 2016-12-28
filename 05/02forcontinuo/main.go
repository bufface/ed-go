package main

import "fmt"

func main() {
	// Sería el equivalente a un "while" en otros lenguajes
	a := 5

	for a > 0 {
		a--
		fmt.Println(a)
	}
}
