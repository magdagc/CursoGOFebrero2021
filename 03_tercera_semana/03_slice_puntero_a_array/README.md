# Slice como un puntero a array

1. En el archivo `main.go` vemos un ejemplo de un array y dos slices que apuntan al mismo array.  
2. El array es inicializado en un principio con los nombres de los Beatles.  
3. Cada slice apunta a un lugar del array inicialmente creado.  
4. Al modificarse un valor del slice `b` (que en realidad lo que modifica es un valor del array `nombres`), tanto el array `nombres` como el slice `a` se ven afectados también ya que todos apuntan al mismo lugar.  
