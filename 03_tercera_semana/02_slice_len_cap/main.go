package main

import "fmt"

func main() {
	fmt.Println("** Creo un slice de enteros de manera literal **")
	s := []int{2, 3, 5, 7, 11, 13}
	imprimirSlice(s)
	fmt.Println("\n** Corto el slice para que tenga longitud cero **")
	s = s[:0]
	imprimirSlice(s)
	fmt.Println("\n** Extiendo su longitud **")
	s = s[:4]
	imprimirSlice(s)
	fmt.Println("\n** Remuevo sus dos primeros valores **")
	s = s[2:]
	imprimirSlice(s)
}
func imprimirSlice(s []int) {
	fmt.Println("\nImprimo la longitud, la capacidad y el valor del slice")
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
