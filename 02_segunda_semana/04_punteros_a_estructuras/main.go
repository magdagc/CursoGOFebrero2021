package main

import "fmt"

type usuario struct {
	titulo   string
	apellido string
}

func main() {
	u := usuario{titulo: "Dr.", apellido: "Jekyll"}

	usuarioValor := u
	usuarioPuntero := &u

	u.titulo = "Mr."
	u.apellido = "Hyde"

	fmt.Printf("Usuario inicializado con el valor de u:\n\tTítulo: %s\n\tApellido:%s\n", usuarioValor.titulo, usuarioValor.apellido)
	fmt.Printf("Usuario inicializado con el puntero de u:\n\tTítulo: %s\n\tApellido:%s\n", usuarioPuntero.titulo, usuarioPuntero.apellido)

}
