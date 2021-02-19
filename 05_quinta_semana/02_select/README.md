# Select

1. En este ejemplo vemos dos canales `c1` y `c2` que reciben señales de tipo entero.  
2. Vamos a lanzar tres goroutines, cada una con un número de mensaje según cuándo disparo la goroutine y un timeout para enviar esa señal cuando se terminó la "tarea".  
3. Vemos que no llegan en orden las señales y en el select se van mostrando a medida que llegan al canal que la tenga disponible.  
4. Esto en un caso real en donde cada tarea tiene su duración, nos permite no depender de tareas que quizás demoran mucho si es que no son necesarias (ahora, si necesitamos de una tarea para ejecutar la siguiente, ya no podremos usar esto porque vemos que se ejecutan en un orden distinto al esperado).  
