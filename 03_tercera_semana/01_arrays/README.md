# Arrays

1. En el ejemplo del archivo `main.go` vemos ejemplos de declaración y uso de arrays.  
2. En el primer caso declaramos una variable cuyo tipo es un array de enteros con 5 posiciones. Al declararlo de esta forma, el contenido de cada lugar de ese array estará dado por el **zero value** del tipo de dato en cada posición (como este ejemplo es un array de `int`, será un array lleno de ceros). Los arrays se imprimen por defecto entre corchetes.  
3. A ese mismo array le pisamos el valor de la posición 4 (la última) e imprimimos cómo queda.  
4. Luego lo que hacemos es obtener un valor en concreto del array accediendo al mismo por índice.  
5. Para obtener la longitud de un array hacemos uso de la función `len(<array>)` que nos provee go.  
6. Otra forma de declarar un array es la forma literal, en donde le indicamos los valores que tendrá ese arreglo. En este caso los valores que le proveo en la declaración son menos que su longitud, así que completará lo que falte con el **zero value** del tipo que sea (al ser un array de `bool`, el **zero value** es un `false`).  
7. Un array podría tener como tipo de dato otro array. En el último caso vemos este "array de dos dimensiones" que podríamos tratar como una matriz de 2x3.  
