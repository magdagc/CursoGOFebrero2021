package main

import (
	"fmt"
	"time"
)

func main() {
	ahora := time.Now()
	var mensaje string

	switch {
	case ahora.Hour() < 12:
		mensaje = "Buenos días"
	case ahora.Hour() < 20:
		mensaje = "Buenas tardes"
	default:
		mensaje = "Buenas noches"
	}

	fmt.Println(mensaje)
}
