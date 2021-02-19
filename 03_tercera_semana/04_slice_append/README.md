# Agregar elementos a un slice

1. En este ejemplo vemos cómo, a través de la función `append`, podemos agregar **n** elementos a un slice y, de ser necesario, redimensionarlo.  
2. Como se crea con una capacidad de 0, apenas agregamos el primer elemento el slice crece tanto en longitud como en capacidad.  
3. En el último caso el slice crece en capacidad más de la longitud que tiene, alocando más espacio en el array intrínseco al prever que podría crecer en longitud luego.  
