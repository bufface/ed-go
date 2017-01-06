package main

import (
	"fmt"

	"Projects/ed_go/06/07paquetes/despedida"
	"Projects/ed_go/06/07paquetes/saludar"
)

func main() {
	nombre := "Cristian"
	saludar.MeVes = "Asigando desde Main.go"

	// saludar.Saludar(nombre)
	fmt.Println(saludar.MeVes)

	despedida.Despedirse(nombre)
}
