package main

import "fmt"

func main() {

	fmt.Println("** For sólo con condición, iteramos de 1 a 3 inclusive **")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("\n** For con inicialización de variable, condición e incremento, iteramos de 7 a 9 inclusive **")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("\n** Iteración sin condición, sigue entrando hasta que haya un corte (en este caso el break que se ejecuta en la primera iteración) **")
	for {
		fmt.Println("iteración sin condición")
		break
	}

	fmt.Println("\n** Iteramos de 0 a 5 inclusive pero sólo imprimimos los números impares **")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
