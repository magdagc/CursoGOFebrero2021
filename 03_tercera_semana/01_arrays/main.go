package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("** Array de enteros de 5 posiciones **\n", a)

	a[4] = 100
	fmt.Println("\n** Asigno 100 a la posición 4 e imprimo el array **\n", a)
	fmt.Println("\n** Obtengo la posición 4 del array **\n", a[4])

	fmt.Println("\n** Longitud del array **\n", len(a))

	b := [5]bool{true, false, true, false}
	fmt.Println("\n** Declaro un array de booleanos de manera literal y lo imprimo **\n", b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("\n** Declaro un array de dos posiciones que en cada posición tiene un array de tres posiciones de enteros **\n", c)
}
