package main

import "fmt"

var poblacion map[string]int

func main() {
	poblacion = make(map[string]int)

	poblacion["Ushuaia"] = 74752
	poblacion["CABA"] = 3075646

	fmt.Println(poblacion)
	fmt.Println(poblacion["Ushuaia"])
}
