package main

import "fmt"

type vertice struct {
	lat, long float64
}

var m = map[string]vertice{
	"Bell Labs": vertice{
		lat: 40.68433, long: -74.39967,
	},
	"Google": vertice{
		lat: 37.42202, long: -122.08408,
	},
}

func main() {
	for clave, valor := range m {
		fmt.Println(clave, valor)
	}
}
