# Interfaz de error

1. En este ejemplo vemos cómo se está usando una estructura cuyo puntero implementa la interfaz `error`.  
2. En el caso que vemos, nuestro `*errorCustom` está entrando siempre al condicional que lo compara con `nil` (parece que siempre devolviera un error), ¿por qué pasa esto? ¿qué estamos devolviendo? (pista: ¿cuál es el zero value de una estructura? ¿cuál es el zero value de un puntero?).  
3. Pensemos cómo podemos devolver un `*errorCustom` que entre al condicional de `err != nil` sólamente cuando sucede un error.  
