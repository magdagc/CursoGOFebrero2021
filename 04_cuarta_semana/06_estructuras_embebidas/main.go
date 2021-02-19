package main

import "fmt"

type motor struct {
	cilindros int
}
type auto struct {
	motor
	marca string
}

func (m motor) sonar() {
	fmt.Printf("Mis %v cilindros hacen poco ruido\n", m.cilindros)
}
func (a auto) andar() {
	fmt.Printf("Con mis %v cilindros ando bien rapido\n", a.cilindros)
}
func main() {
	fitito := auto{motor: motor{cilindros: 4}, marca: "Fiat"}
	fitito.sonar()
	fitito.andar()
}
