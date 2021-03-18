# Abrir conexión a una base de datos

1. En este ejemplo utilizamos una base en memoria llamada [RamSQL](https://github.com/proullon/ramsql)  
2. Si quisiéramos modificar el motor de base de datos, deberíamos usar un driver distinto.  
3. Los drivers disponibles pueden encontrarlos en [esta página](https://github.com/golang/go/wiki/SQLDrivers).
4. La función `sql.Open()` devuelve un puntero a la estructura `DB` dentro del paquete `sql` y un `error`.
5. Podemos pensar que ese `error` puede ser algo distinto de `nil` cuando no podemos lograr conectarnos a la base, pero no es tan así. La función realmente devuelve un error cuando la url de conexión tiene mal su formato o el driver no es reconocido por ejemplo (probemos quitar el import de ramsql).
6. Para comprobar si se pudo conectar efectivamente a la base de datos, podemos usar el método `Ping()` que implementa el puntero a la estructura `DB`.  
