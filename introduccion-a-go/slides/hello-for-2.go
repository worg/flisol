package main

import "fmt"

func main() {
	días := []string{`Lunes`, `Martes`, `Miércoles`, `Jueves`, `Viernes`, `Sábado`, `Domingo`}
	// Omitimos el primer elemento, que es el key [índice] de nuestro slice
	for _, dia := range días {
		fmt.Println(`Hoy es:`, dia)
	}
}
