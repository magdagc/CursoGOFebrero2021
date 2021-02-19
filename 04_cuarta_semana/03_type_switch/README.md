# Type switch

1. En este ejemplo vemos cómo usar el `switch` para ver de qué tipo es una variable (recordemos que una variable tiene **un solo tipo**, no puede ser float32 y float64 al mismo tiempo por ejemplo).  
2. Notemos cómo, cuando le pasamos un puntero a una variable `int`, no lo reconoce como `int`.  
3. Si queremos que caiga por algún case, deberíamos agregar el caso en donde el tipo coincida con `*int`, hagámoslo.  
