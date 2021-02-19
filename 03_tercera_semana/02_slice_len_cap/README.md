# Longitud y capacidad para un slice

1. En el archivo `main.go` podemos ver cómo varían la longitud y la capacidad de un slice a medida que vamos modificando a qué parte del array intrínseco apunta.  
2. En el primer caso vemos que tanto la longitud como la capacidad están en 6 (se usan los 6 elementos alocados del array al que apuntamos).  
3. Luego recortamos ese slice para que no tenga ningún valor del array intrínseco. Los elementos siguen existiendo en el array y la capacidad es la misma, pero en nuestro slice estamos apuntando a una sección que no tiene elementos.  
4. Luego extendemos la longitud del slice, lo que hace que apunte a los primeros cuatro elementos del array intrínseco. Aquí se modifica su longitud pero no la capacidad.  
5. En el último ejemplo recortamos los dos primeros elementos del slice. En este caso sí se modifica la capacidad porque ya no puedo acceder a esos dos primeros elementos.  
