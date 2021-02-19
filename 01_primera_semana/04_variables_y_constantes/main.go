package main

import "fmt"

// Booleanos sin inicializar
var c, python, java bool

// Enteros inicializados en 1 y 2 correspondientemente
var uno, dos int = 1, 2

// Constante
// ¿Qué sucede si intentamos modificar el valor más tarde en el código?
const pi float32 = 3.14

func main() {
	// Entero sin inicializar
	var i int

	// No defino un tipo para las variables, directamente las inicializo y el tipo se infiere
	var verdadero, falso, no = true, false, "no!"

	// Sin usar var ni definirle un tipo, se infiere
	k := 3

	fmt.Printf("Booleanos sin inicializar:\n c: %v\n python: %v\n java: %v\n", c, python, java)
	fmt.Printf("Enteros inicializados:\n uno: %d\n dos: %d\n", uno, dos)
	fmt.Printf("Constante:\n pi con formato float por defecto:%f\n pi con 1 decimal de precisión: %.1f\n pi con 3 decimales de precisión: %.3f\n", pi, pi, pi)
	fmt.Printf("Entero sin inicializar:\n i: %d\n", i)
	fmt.Printf("Variables con inferencia de tipo:\n verdadero: %v\n falso: %v\n no: %s\n k: %d\n", verdadero, falso, no, k)
}
