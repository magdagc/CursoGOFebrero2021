package main

import "fmt"

func main() {

	m := make(map[int]string)
	m[1] = "uno"
	m[2] = "dos"
	m[100] = "cien"

	fmt.Println("Mapa:", m)

	v1 := m[1]
	fmt.Println("v1: ", v1)

	fmt.Println("Longitud: ", len(m))

	_, existe := m[2]
	fmt.Println("Existe antes del delete:", existe)

	delete(m, 2)

	fmt.Println("Mapa:", m)

	_, existe = m[2]
	fmt.Println("Existe después del delete:", existe)

}
