# WaitGroup, atomic y Mutex

1. En este caso vemos por un lado el `sync.WaitGroup` al cual le indicamos que espere a que terminen las 150 goroutines que llamamos (se ve en la llamada a `wg.Add(3)` que se realiza las 50 veces que pasa por la iteración).  
2. Por otro lado, vemos que hay tres funciones diferentes a las que podemos llamar para incrementar un contador mil veces, una de ellas utiliza el paquete `atomic`, otra utiliza la estructura `sync.Mutex` y la otra no utiliza nada de eso.  
3. Cuando se invoca a la función `aumentarContador` como goroutine, no podemos asegurar que el contador quede sincronizado (ejecutarlo varias veces para comprobarlo, a veces puede que devuelva 50000).  
4. Cuando se invoca la función `aumentarContadorConAtomic` sí se puede asegurar porque el incremento lo hacemos con la función `atomic.AddUint64`.  
5. Otra forma de asegurarlo es utilizando `sync.Mutex` que bloqueará los recursos compartidos (en este caso el entero que incrementamos) hasta desbloquearlo explícitamente.  
