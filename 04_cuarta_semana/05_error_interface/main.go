package main

import (
	"fmt"
)

type errorCustom struct {
	mensaje string
}

func (ec *errorCustom) Error() string {
	return ec.mensaje
}

func main() {

	err := devuelvoUnErrorCustom(true)

	if err != nil {
		fmt.Printf("Error custom pasando true: %s\n", err.Error())
	}

	err = devuelvoUnErrorCustom(false)

	if err != nil {
		fmt.Printf("Error custom pasando false: %s\n", err.Error())
	}

}

func devuelvoUnErrorCustom(debeFallar bool) error {
	var err errorCustom

	if debeFallar {
		err.mensaje = "soy un error"
	}

	return &err

}
