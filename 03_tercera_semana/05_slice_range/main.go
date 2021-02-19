package main

import "fmt"

var potencias = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for indice, valor := range potencias {
		fmt.Printf("2**%d = %d\n", indice, valor)
	}
	for _, valor := range potencias {
		fmt.Printf("%d\n", valor)
	}
}
