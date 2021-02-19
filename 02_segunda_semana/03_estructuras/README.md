# Estructuras  

1. En el archivo `dominio/vertice.go` vemos un ejemplo de la estructura `Vertice` que tiene dos campos (`X` e `Y`, ambos enteros).  
2. Hacemos uso de esa estructura en el archivo `main` creando un nuevo `Vertice` llamado `v` inicializado con valores concretos.  
3. Luego pisamos el valor que tenía `v` en `X` e imprimimos `v`.  
4. La forma de inicializar una estructura que vemos allí es sin indicarle explícitamente cuál es el valor de `X` ni cual es el valor de `Y`, lo hace en orden.  
5. Se puede especificar cuál es el valor de cada campo y no necesariamente seguir el orden de la estructura como fue definida en el `type`, es decir, se podría haber inicializado de la misma forma como `Vertice{Y: 2, X: 1}`.  
6. También se puede omitir algún campo de la estructura, en ese caso quedaremos con los campos que no se definieron en su **zero value**, que para los `int` es cero.  
7. Tanto la estructura `Vertice` como sus campos están exportados porque empiezan con mayúscula, ¿qué pasaría si esto no fuera así para `Vertice`? ¿y para el campo `X`? ¿y para el campo `Y`?  
