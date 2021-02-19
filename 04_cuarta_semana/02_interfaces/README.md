# Interfaces

1. En este ejemplo vemos cómo el puntero a `Implementador` implementa la interfaz `fmt.Stringer` y la interfaz que creamos `dominio.StringerFalso`.  
2. ¿Qué pasa si en lugar de pasar un puntero a `Implementador` pasamos la estructura `Implementador`?  
3. ¿Qué pasa si intentamos pasar un type que no implemente la interfaz `dominio.StringerFalso`? Probemos eliminar el método `StringAlReves` definido en `implementador.go`.  
