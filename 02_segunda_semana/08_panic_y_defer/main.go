package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Salí de la función main")
	f := crearArchivo("./03_estructuras_de_control/04_panic_y_defer/defer.txt")
	//f := crearArchivo("./directorio_que_no_existe/defer.txt")
	defer cerrarArchivo(f)
	escribirArchivo(f)
}

func crearArchivo(p string) *os.File {
	fmt.Println("Creando el archivo...")
	//defer fmt.Println("Salí de la función crearArchivo")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func escribirArchivo(f *os.File) {
	fmt.Println("Escribiendo en el archivo...")
	//defer fmt.Println("Salí de la función escribirArchivo")
	fmt.Fprintln(f, "hola, soy el contenido de un archivo")

}

func cerrarArchivo(f *os.File) {
	fmt.Println("Cerrando el archivo...")
	//defer fmt.Println("Salí de la función cerrarArchivo")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
