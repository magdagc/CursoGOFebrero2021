package main

import "fmt"

func main() {
	f()
	fmt.Println("Salí bien de f")
}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Me recupero en f", r)
		}
	}()
	fmt.Println("Invoco a g")
	g(0)
	fmt.Println("Salí bien de g")
}
func g(i int) {
	if i > 3 {
		fmt.Println("¡Estoy entrando en pánico!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer en g", i)
	fmt.Println("Hola, estoy en g", i)
	g(i + 1)
}
