package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var contador, contadorAtomic, contadorMutex uint64

	var wg sync.WaitGroup
	m := sync.Mutex{}

	for i := 0; i < 50; i++ {
		wg.Add(3)

		go aumentarContador(&contador, &wg)
		go aumentarContadorConAtomic(&contadorAtomic, &wg)
		go aumentarContadorConMutex(&contadorMutex, &wg, &m)
	}

	wg.Wait()

	fmt.Println("contador:", contador)
	fmt.Println("contadorAtomic:", contadorAtomic)
	fmt.Println("contadorMutex:", contadorMutex)
}

func aumentarContador(contador *uint64, wg *sync.WaitGroup) {
	for c := 0; c < 1000; c++ {
		*contador++
	}
	wg.Done()
}

func aumentarContadorConAtomic(contador *uint64, wg *sync.WaitGroup) {
	for c := 0; c < 1000; c++ {
		atomic.AddUint64(contador, 1)
	}
	wg.Done()
}

func aumentarContadorConMutex(contador *uint64, wg *sync.WaitGroup, m *sync.Mutex) {
	for c := 0; c < 1000; c++ {
		m.Lock()
		*contador++
		m.Unlock()
	}
	wg.Done()
}
