package main

import "fmt"

func main() {
	fmt.Println("** Creo un array de strings con los nombres de los Beatles **")
	nombres := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(nombres)

	fmt.Println("\n** Creo un slice con los dos primeros nombres (lo llamaremos \"a\") **")
	a := nombres[0:2]
	fmt.Println("** Creo un slice con el segundo y tercer nombre (lo llamaremos \"b\") **")
	b := nombres[1:3]
	fmt.Println("\n** Imprimo el slice a y el slice b **")
	fmt.Println(a, b)
	fmt.Println("\n** Modifico el primer nombre del slice b **")
	b[0] = "XXX"
	fmt.Println("\n** Imprimo el slice a y el slice b **")
	fmt.Println(a, b)
	fmt.Println("\n** Imprimo el array inicial de nombres **")
	fmt.Println(nombres)
}
