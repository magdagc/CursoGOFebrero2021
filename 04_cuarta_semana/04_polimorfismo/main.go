package main

import (
	"fmt"

	"github.com/magdagc/CursoGOFebrero2021/05_metodos_interfaces/04_polimorfismo/animal"
)

// Animal es una interfaz para que implementen animales
// y emitan el sonido que les corresponda
type Animal interface {
	EmitirSonido() string
}

func main() {
	var animales []Animal
	gato := &animal.Gato{}
	perro := &animal.Perro{}
	animales = append(animales, gato, perro)
	for _, animal := range animales {
		fmt.Println(animal.EmitirSonido())
	}
}
