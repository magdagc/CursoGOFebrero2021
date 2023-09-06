package main

import (
	helloworld "github.com/magdagc/CursoGOFebrero2021/01_primera_semana/01_hello_world"
)

func main() {
	// Hago referencia a la función ImprimirHelloWorld de mi paquete helloworld
	helloworld.ImprimirHelloWorld()

	// ¿Qué pasa en este caso?
	//helloworld.imprimirHelloWorld()
}
