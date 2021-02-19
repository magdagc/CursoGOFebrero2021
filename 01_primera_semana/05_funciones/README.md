# Funciones

1. Las funciones tienen la siguiente sintaxis en go por defecto (según la cantidad de variables de entrada y salida vemos que puede ir variando):
`func <nombre>(<entradas>) <salidas>`
   1. Ejemplo de una función con un entero de entrada y un float32 de salida:  
    `func miFuncion(x int) float32`
   2. Ejemplo de una función con un entero de entrada y un float32 de salida con nombre:  
    `func miFuncion(x int) (y float32)`
   3. Ejemplo de una función con dos enteros de entrada, un float32 y dos strings de salida (se puede escribir de cualquiera de estas dos formas que son equivalentes):  
    `func miFuncion(x, y int) (float32, string, string)`  
    `func miFuncion(x int, y int) (float32, string, string)`  
   4. Ejemplo de una función sin entradas y con un string de salida:  
    `func miFuncion() string`
   5. Ejemplo de una función con un string de entrada y sin salida:  
    `func miFuncion(s string)`  
2. En el archivo `main.go` se ejecutan cuatro funciones diferentes (ninguna está exportada porque las usamos dentro del mismo paquete en donde fueron declaradas).
   1. La función `sumar` que recibe dos enteros como entrada y devuelve un entero con la suma.  
   2. La función `restar` que recibe dos enteros como entrada y devuelve un entero con la resta.  
   3. La función `sumarMuchosNumeros` que recibe **n** enteros como entrada y devuelve un entero con la suma.  
   4. La función `imprimirMensaje` que recibe tres enteros como entrada y una cadena de caracteres y no tiene salida.  
   5. La función `transformarCadena` que no recibe entradas y devuelve tres funciones que convierten un string a determinados formatos.  
3. Podríamos no querer usar todas las variables de salida en algún caso (por ejemplo en la función `transformarCadena`), en esos casos usamos el guión bajo `_` en lugar del nombre de variables que no queremos usar. Imaginemos que sólo queremos devolver la función que devuelve la cadena toda en mayúsculas, ¿qué haríamos?
