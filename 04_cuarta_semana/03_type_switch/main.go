package main

import "fmt"

func main() {
	i := 24
	s := "holis"
	f := 3.16

	evaluarTipo(i)
	evaluarTipo(s)
	evaluarTipo(f)
	evaluarTipo(&i)
}

func evaluarTipo(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Es un entero")
	case string:
		fmt.Println("Es un string")
	default:
		fmt.Println("No conozco ese tipo")
	}
}
