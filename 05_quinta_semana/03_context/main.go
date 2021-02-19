package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("** CONTEXTO CON CANCELACIÓN MANUAL DESPUÉS DE LOS 3 SEGUNDOS **")
	probarContextoConCancelacion()
	fmt.Println("** CONTEXTO CON TIMEOUT DE 4 SEGUNDOS**")
	probarContextoConTimeOut()
	fmt.Println("** CONTEXTO CON DEADLINE 10 SEGUNDOS DESPUÉS DE LA FECHA ACTUAL**")
	probarContextoConDeadline()
}

type claveContext string

const claveTipoContext = claveContext("tipo-context")

func probarContextoConCancelacion() {
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	go task(context.WithValue(cancelCtx, claveTipoContext, "cancel"))

	time.Sleep(time.Second * 3)
	fmt.Println("Cancelo manualmente")
	cancelFunc()
	time.Sleep(time.Second * 1)
}

func probarContextoConTimeOut() {
	ctx := context.Background()
	timeoutCtx, cancelTimeoutFunc := context.WithTimeout(ctx, 4*time.Second)
	defer cancelTimeoutFunc()

	go task(context.WithValue(timeoutCtx, claveTipoContext, "timeout"))

	time.Sleep(time.Second * 5)
}

func probarContextoConDeadline() {
	ctx := context.Background()
	fechaActual := time.Now()
	fechaDeadline := fechaActual.Add(10 * time.Second)

	fmt.Println("Fecha actual: " + fechaActual.Format("2006-01-02 15:04:05"))
	fmt.Println("Fecha deadline: " + fechaDeadline.Format("2006-01-02 15:04:05"))

	deadlineCtx, cancelDeadlineFunc := context.WithDeadline(ctx, fechaDeadline)
	defer cancelDeadlineFunc()

	go task(context.WithValue(deadlineCtx, claveTipoContext, "deadline"))

	time.Sleep(time.Second * 11)
}

func task(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Salgo de manera elegante")
			fmt.Printf("El Value configurado en tipo-context es: %v\n", ctx.Value(claveTipoContext))
			fmt.Printf("Lo que devuelve Err para este contexto es: %v\n\n\n", ctx.Err())
			return
		default:
			fmt.Printf("Tiempo transcurrido en segundos: %d\n", i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
