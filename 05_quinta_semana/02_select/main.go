package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan int), make(chan int)
	go enviarSenialDespuesDeTimeout(3, 1, c1)
	go enviarSenialDespuesDeTimeout(2, 2, c2)
	go enviarSenialDespuesDeTimeout(1, 3, c1)
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("recibido", msg1)
		case msg2 := <-c2:
			fmt.Println("recibido", msg2)
		}
	}
}

func enviarSenialDespuesDeTimeout(i time.Duration, numeroMensaje int, c chan int) {
	time.Sleep(time.Duration(i) * time.Second)
	c <- numeroMensaje
}
