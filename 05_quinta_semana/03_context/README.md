# Paquete context

1. Aquí vemos tres ejemplos de cómo usar context para:  
   1. Enviar una cancelación manualmente.  
   2. Enviar una cancelación después de una duración determinada.  
   3. Enviar una cancelación a partir de una fecha límite.  
2. En cada contexto aprovechamos para mostrar cómo configurar un Value para la clave "tipo-context".  
3. Cuando utilizamos la función `Background()` del paquete `context`, se nos provee con una variable de un tipo que cumple con la interfaz `Context`.  
4. Los contextos que utilizamos son derivados de ese contexto vacío y vamos indicándole en cada caso cómo queremos que se comporte en cuanto a la cancelación.  
