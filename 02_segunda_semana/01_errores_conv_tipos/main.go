package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Esta función del paquete strconv convierte un string a un int
	// El entero queda en la variable i
	// Si no se pudiera convertir, la variable err contendría un error
	i, errAtoi := strconv.Atoi("42")

	if errAtoi != nil {
		fmt.Printf("No se puede convertir el numero: %v\n", errAtoi)
		return
	}

	fmt.Printf("El valor convertido a entero es %d\n", i)

	division, errDividir := dividir(8, 0)

	if errDividir != nil {
		fmt.Printf("No se pudo ejecutar la división: %v\n", errDividir.Error())
	} else {
		fmt.Printf("El resultado de la división es %f", division)
	}
}

func dividir(dividendo, divisor int) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("no se puede usar %d como divisor", divisor)
		// return 0, errors.New("no se puede dividir por cero")
	}
	return (float64(dividendo) / float64(divisor)), nil
}
