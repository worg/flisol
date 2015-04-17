package main

import "fmt"

// Variables globales
var (
	a int                // Definida por tipo
	b = `Soy una cadena` // Definida por su valor
)

func main() {
	c := a

	fmt.Printf("a: %v \nb: %v \nc: %v", a, b, c)
}
