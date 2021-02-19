# Goroutines

1. En este ejemplo vemos distintas llamadas a una misma función (una llamada directa y dos como goroutine) y una llamada a una función anónima como goroutine.  
2. Ejecutando el main como está, vemos que la primera goroutine no llega a finalizar porque la demora que tiene entre mensaje y mensaje hace que termine después del final de la función main.  
3. La ejecución no va a esperar a que se terminen todas las goroutines, por eso es **muy importante** que si llamamos a alguna función como goroutine seamos conscientes de esto.  
4. Hay maneras de asegurarse que finalicen las goroutines, pero siempre que llamemos a alguna función como goroutine tenemos que tener esto en cuenta porque podríamos pensar que algo se ejecuta y no lo está haciendo.  
5. Por otro lado, ¿qué sucede si sacamos el `go` de la llamada que dice "primera goroutine"? Probémoslo.  
