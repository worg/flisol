package main

import "fmt"

// Constantes Globales
const (
	_      = iota // iota [autoincremento]
	KB int = 1 << (10 * iota)
	MB
	GB
	sufijo = `bytes` // también pueden ser cadenas
)

func main() {
	fmt.Printf("KB: %d %s\n", KB, sufijo)
	fmt.Printf("MB: %d KB ó %d bytes\n", MB/KB, MB)
	fmt.Printf("GB: %d MB  ó %d bytes\n", GB/MB, GB)
}
