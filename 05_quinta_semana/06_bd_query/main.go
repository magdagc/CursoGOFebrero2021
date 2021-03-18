package main

import (
	"database/sql"
	"fmt"

	// Necesitamos importar este paquete aunque no
	// se use explícitamente
	// ¿Qué pasa si no lo importamos?
	_ "github.com/proullon/ramsql/driver"
)

type userAddress struct {
	usuarioID   int64
	direcciones []address
}

type address struct {
	calle  sql.NullString
	numero sql.NullInt32
}

func main() {

	// Abrimos la conexión
	db, err := sql.Open("ramsql", "PruebaQuery")
	if err != nil {
		fmt.Printf("sql.Open : Error : %s\n", err.Error())
		return
	}
	defer db.Close()

	// Cargamos datos en la base a la que estamos conectades
	err = cargarDatos(db)

	if err != nil {
		fmt.Printf("Error en la carga de datos : %s\n", err.Error())
		return
	}

	// Buscamos datos de lo que ya cargamos
	usuarioDireccion, err := buscarUsuarioDireccion(db, 1)

	if err != nil {
		fmt.Printf("Error en la búsqueda de datos : %s\n", err.Error())
		return
	}

	fmt.Println(usuarioDireccion)

}

func cargarDatos(db *sql.DB) error {
	batch := []string{
		`CREATE TABLE address (id BIGSERIAL PRIMARY KEY, street TEXT, street_number INT);`,
		`CREATE TABLE user_addresses (address_id INT, user_id INT);`,
		`INSERT INTO address (street, street_number) VALUES ('rue Victor Hugo', 32);`,
		`INSERT INTO address (street, street_number) VALUES ('boulevard de la République', 23);`,
		`INSERT INTO address (street, street_number) VALUES ('rue Charles Martel', 5);`,
		`INSERT INTO address (street, street_number) VALUES ('chemin du bout du monde ', 323);`,
		`INSERT INTO address (street, street_number) VALUES ('boulevard de la liberté', 2);`,
		`INSERT INTO address (street, street_number) VALUES ('avenue des champs', 12);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 1);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 1);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 2);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (2, 3);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 4);`,
		`INSERT INTO user_addresses (address_id, user_id) VALUES (4, 5);`,
	}

	for _, b := range batch {
		_, err := db.Exec(b)
		if err != nil {
			return fmt.Errorf("sql.Exec: Error: %s\n", err)
		}
	}

	return nil
}

func buscarUsuarioDireccion(db *sql.DB, usuarioID int) (userAddress, error) {

	usuarioDireccionBuscado := userAddress{usuarioID: int64(usuarioID)}

	// Los parámetros se escriben diferente según cada driver
	// Para este driver usamos la sintaxis $1, $2, $3...
	// Cuando llamamos al método Query pasándole parámetros reemplaza eso mismo
	query := `SELECT address.street_number, address.street FROM address 
							JOIN user_addresses ON address.id=user_addresses.address_id 
							WHERE user_addresses.user_id = $1;`
	rows, err := db.Query(query, usuarioID)

	if err != nil {
		return usuarioDireccionBuscado, fmt.Errorf("db.Query : Error : %s\n", err)
	}

	for rows.Next() {
		direccionAAgregar := address{}
		rows.Scan(&direccionAAgregar.numero, &direccionAAgregar.calle)
		usuarioDireccionBuscado.direcciones = append(usuarioDireccionBuscado.direcciones, direccionAAgregar)
	}

	return usuarioDireccionBuscado, nil
}
