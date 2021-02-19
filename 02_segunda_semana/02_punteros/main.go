package main

import "fmt"

func main() {
	i := 42

	p := &i // puntero a i
	fmt.Println("Imprimo el valor al que apunta el puntero p accediendo al mismo como *p")
	fmt.Println(*p) // accedo a i a través del puntero
	*p = 21         // modifico i a través del puntero
	fmt.Println("Imprimo el valor de i")
	fmt.Println(i) // i fue modificado
}
