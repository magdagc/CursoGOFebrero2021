package main

import (
	"fmt"
	"time"
)

func f(from string, duracion time.Duration) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(duracion)
	}
}

func main() {

	go f("primera goroutine", 200*time.Millisecond)

	go f("segunda goroutine", 50*time.Millisecond)

	go func(msg string) {
		fmt.Println(msg)
	}("una goroutine anónima")

	f("llamada directa", time.Microsecond)

	time.Sleep(time.Second)
	fmt.Println("listo")
}
