# Manejo de errores  

1. En go no existen las excepciones, trabajamos con errores. Utilizamos generalmente la posibilidad de devolver varias salidas en una función para devolver como última salida un error si es que esa función puede fallar (el error no es un tipo de dato sino una interfaz, veremos estos conceptos más tarde).  
2. Comprobamos que una función salió sin errores al comparar el error que sale con `nil`, si el error no es igual a `nil`, es un patrón común que manejemos ese error y hagamos un return de la función.  
3. Dos formas de crear errores para devolver en nuestras funciones son `fmt.Error("mensaje")` o `errors.New("mensaje")`.  
4. En el archivo `main.go` vemos ejemplos de los puntos 2 y 3.  
5. Si quisiéramos devolver únicamente el valor de la función sin manejar el error que devuelve, usaríamos el guión bajo, probemos hacer eso con la función `strconv.Atoi("42")` que sabemos que convierte bien el "42" en un entero.  
6. El caso de la función `dividir(dividendo, divisor int) (float64, error)` es una función no exportada que devuelve un error además del resultado cuando se usa un divisor igual a cero.  
7. Lo que vemos en el caso feliz de la división es una conversión de tipos, en donde los `int` se convierten a `float64`, la sintaxis es `<tipo_convertido>(<variable_con_tipo_a_convertir>)`, en este caso `float64(variableEntera)` que nos da un `float64`.
