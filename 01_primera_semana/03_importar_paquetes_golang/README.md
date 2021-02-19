# Usemos los paquetes de golang

En este ejemplo usamos el paquete rand que golang trae (y el fmt que ya habíamos usado).  

1. En el archivo 01_package_rand.go vemos la función imprimirNumeroRand que imprime un número que a primera vista parece ser aleatorio.  
2. Ejecutemos desde el directorio del proyecto:  
`cd 01_conceptos_basicos/03_importar_paquetes_golang`  
`go run main.go`  
3. Ese número no es realmente aleatorio. Para lograr obtener valores distintos, podemos agregar justo antes de imprimir el número la siguiente línea:  
`rand.Seed(time.Now().UnixNano())`  
4. Esto configurará como semilla el tiempo actual en milisegundos para darnos valores distintos. [La documentación del paquete rand](https://golang.org/pkg/math/rand/) nos muestra ejemplos de cómo usar ese paquete, es lo que deberíamos consultar con los paquetes que importemos.  
