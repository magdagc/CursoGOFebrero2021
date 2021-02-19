package main

import (
	"fmt"
	"strings"
)

func main() {

	var a, b, c, d int = 8, 3, 5, 92

	suma := sumar(a, b)
	resta := restar(a, b)

	fmt.Printf("Suma:\n %d + %d = %d\n", a, b, suma)
	fmt.Printf("Resta:\n %d - %d = %d\n", a, b, resta)

	sumaMuchosNumeros := sumarMuchosNumeros(a, b, c, d)

	fmt.Printf("Suma:\n %d + %d + %d + %d = %d\n", a, b, c, d, sumaMuchosNumeros)

	sumaMuchosNumeros = sumarMuchosNumeros(d)

	fmt.Printf("Sumo un solo número:\n %d = %d\n", d, sumaMuchosNumeros)

	sumaMuchosNumeros = sumarMuchosNumeros()

	fmt.Printf("Llamo a sumaMuchosNumeros sin parámetros:\n %d\n", sumaMuchosNumeros)

	imprimirMensajes(74, 1, "hola", -5)

	s := "CaDeNa De EjEmPlO"

	minusculas, mayusculas, titulo := transformarCadena()

	fmt.Printf("Transformación de cadena de caracteres:\n"+
		" Cadena original: %s\n Cadena en minúsculas: %s\n"+
		" Cadena en mayúsculas: %s\n Cadena como título: %s\n",
		s, minusculas(s), mayusculas(s), titulo(s))

}

func sumar(x, y int) int {
	return x + y
}

func sumarMuchosNumeros(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func restar(x int, y int) (resta int) {
	resta = x - y
	return
}

func imprimirMensajes(m, n int, s string, o int) {
	fmt.Printf("La función imprimirMensajes no tiene salida\n")
	fmt.Printf("Tiene tres entradas numéricas:\n m: %d\n n: %d\n o: %d\n", m, n, o)
	fmt.Printf("Y tiene una entrada de cadena de caracteres:\n s: %s\n", s)
}

func transformarCadena() (func(string) string, func(string) string, func(string) string) {
	minusculas := strings.ToLower
	mayusculas := strings.ToUpper
	titulo := func(s string) string { return strings.Title(strings.ToLower(s)) }

	return minusculas, mayusculas, titulo
}
