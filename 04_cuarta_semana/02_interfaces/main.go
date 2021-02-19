package main

import (
	"fmt"

	"github.com/magdagc/CursoGOFebrero2021/05_metodos_interfaces/02_interfaces/dominio"
)

func main() {
	e := &dominio.Implementador{A: "a", B: "b"}
	ImprimirStringDeStringer(e)
	ImprimirStringsDeStringerFalso(e)
}

// ImprimirStringDeStringer imprime lo que devuelva el método String
// de la entrada
func ImprimirStringDeStringer(s fmt.Stringer) {
	fmt.Println(s.String())
}

// ImprimirStringsDeStringerFalso imprime lo que devuelva el método String
// y el método StringAlReves de la entrada
func ImprimirStringsDeStringerFalso(s dominio.StringerFalso) {
	fmt.Println(s.String())
	fmt.Println(s.StringAlReves())
}
