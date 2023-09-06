package main

import (
	// Paquete de golang, no necesitamos manejo de dependencias,
	// basta con importarlo así
	"fmt"
	"math/rand"
)

func main() {
	// Llamamos a una función dentro del mismo paquete, no necesita ser exportada
	imprimirNumeroRand()
}

// imprimirNumeroRand utiliza el paquete rand para imprimir un número.
func imprimirNumeroRand() {
	fmt.Printf("Mi número favorito es %d\n", rand.Intn(10))
}
