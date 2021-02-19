package main

import (
	helloworld "github.com/magdagc/CursoGOFebrero2021/01_conceptos_basicos/01_hello_world"
)

func main() {
	// Hago referencia a la función ImprimirHelloWorld de mi paquete helloworld
	helloworld.ImprimirHelloWorld()

	// ¿Qué pasa en este caso?
	//helloworld.imprimirHelloWorld()
}
