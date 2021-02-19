# Punteros  

1. En `main.go` vemos un ejemplo de cómo se define un puntero a una variable.  
2. Primero se inicializa una variable `i` que es un entero con un valor inicial.  
3. Luego definimos una variable `p` que será un puntero y estará apuntando a lo que hay en `i`. Para obtener el puntero que hará referencia a lo que haya en `i` se usa el ampersand, así `&i` es un puntero a la variable `i`.  
4. Teniendo el puntero `p` (olvidándonos por un momento que tenemos la variable `i` a nuestra disposición), para acceder o modificar el valor de la variable a la cual apunta debemos anteponer un asterisco al puntero. Así es que asignando un valor a `*p`, realmente se está modificando la variable a la cual apunta (que en este caso es `i`), pasa lo mismo cuando imprimimos `*p`.  
