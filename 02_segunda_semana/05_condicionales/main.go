package main

import (
	"fmt"
)

type persona struct {
	tuvoContactoConCovid bool
	tieneDolorDeCabeza   bool
	tieneTos             bool
	temperatura          float32
}

func main() {

	if sospechosaDeCovid := llenarDatosPersona(); sospechosaDeCovid.tuvoContactoConCovid ||
		sospechosaDeCovid.tieneDolorDeCabeza ||
		sospechosaDeCovid.tieneTos ||
		sospechosaDeCovid.temperatura >= 37.5 {
		fmt.Println("Hacemos hisopado")
	} else {
		fmt.Println("Por ahora no hacemos nada")
	}

}

func llenarDatosPersona() persona {
	return persona{
		tuvoContactoConCovid: false,
		tieneDolorDeCabeza:   true,
		tieneTos:             true,
		temperatura:          37,
	}
}
