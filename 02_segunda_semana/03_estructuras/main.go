package main

import (
	"fmt"

	"github.com/magdagc/CursoGOFebrero2021/02_estructuras_punteros/02_estructuras/dominio"
)

func main() {
	v := dominio.Vertice{1, 2}
	// es lo mismo que decir
	// v := dominio.Vertice{X: 1, Y: 2}
	fmt.Println("Imprimimos v que fue inicializado como vertice{1, 2}")

	v.X = 4
	fmt.Println(v)
}
