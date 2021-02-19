# Panic y defer

1. En el ejemplo de `main.go` vemos cómo, en el camino feliz, se crea un archivo en un directorio existente y se cierra haciendo uso del defer.  
2. Si descomentamos la línea `f := crearArchivo("./directorio_que_no_existe/defer.txt")`, veremos cómo, al llamar a `crearArchivo`, la línea `f, err := os.Create(p)` devolverá un error y llamaremos a la función `panic`.
3. Podemos ver que en este caso no se ejecuta el `defer` que cerraba el archivo declarado en la función `main`, ¿por qué sucede eso?
4. Si queremos tener más noción (tanto en el caso feliz como en el otro) de qué `defer` ejecuta, podemos descomentar las líneas con `defer` que imprimen cuando sale de una función.  
