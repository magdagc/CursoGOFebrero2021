package main

import (
	"database/sql"
	"fmt"

	// Necesitamos importar este paquete aunque no
	// se use explícitamente
	// ¿Qué pasa si no lo importamos?
	_ "github.com/proullon/ramsql/driver"
)

func main() {
	db, err := sql.Open("ramsql", "PruebaConexion")
	if err != nil {
		fmt.Printf("sql.Open : Error : %s\n", err.Error())
	} else {
		fmt.Println("Tenemos una conexión a la base definida")
	}

	// ¿Qué pasa si cerramos la conexión antes de invocar a Ping?
	// db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Printf("db.Ping : Error : %s\n", err.Error())
	} else {
		fmt.Println("¡Nos conectamos correctamente a la base!")
	}

	defer db.Close()
}
