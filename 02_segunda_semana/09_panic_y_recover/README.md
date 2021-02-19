# Panic y recover

1. En el archivo `main.go` vemos un ejemplo más complejo de `recover`.  
2. Dentro de la función `main` se **ejecuta** una llamada a la función `f`.  
3. Lo primero que se hace dentro de la función `f` es **apilar en las llamadas diferidas** una función anónima que ejecutará el `recover` si es que hay algún `panic`. Como es una llamada diferida, no se ejecuta sino que se dejará pendiente para ejecutar una vez que salga de la función.  
4. Lo que se hace luego es **ejecutar** una impresión por pantalla indicando que se invoca a la función `g`.  
5. Dentro de `g` (al cual llamo con un cero en la entrada), **cuando la entrada sea mayor a 3**, **ejecuta** una impresión por pantalla indicando que va a invocar a `panic` y luego **ejecuta** una llamada a `panic`. En un principio no pasará por aquí.
6. Si la entrada aún no llegó a 4 (porque de lo contrario sale por el `panic`)
   1. **Apila en las llamadas diferidas** una impresión por pantalla indicando que es un defer en `g` y el número que fue pasado como entrada en ese momento.  
   2. **Ejecuta** una impresión por pantalla indicando que está en `g` y el número que fue pasado como entrada.  
   3. Vuelve a entrar en `g` recursivamente.  
7. Volviendo a la función `f`, se **ejecuta** una impresión por pantalla indicando que salió de `g`.
8. Cuando sale de la función `f` se **ejecuta** una impresión por pantalla indicando que salió de `f`.  
9. ¿Por qué vemos que no se imprime `Salí bien de g` pero sí se imprime `Salí bien de f`?  
10. Probemos qué pasa si no tenemos `recover` (eliminemos todo el defer de la función `f`).  
