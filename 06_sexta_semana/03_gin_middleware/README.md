# Middlewares con gin

1. En este ejemplo vemos cómo utilizar el framework de gin para incluir un log justo antes de cada request que se envía.  
2. El handler registrado es para la URL `localhost:8080` con método GET, si ejecutamos este request, veremos que desde la terminal que (recomendamos para enviar el request empezar a usar Postman para acostumbrarnos a la herramienta, pero se puede enviar por otros medios).  
3. También pueden intentar enviar otro request distinto que nos dé 404, veremos que en esos casos también se loguea (por ejemplo, hacer un POST en lugar de un GET con la misma URL).  
4. En la collection de postman **api_rest.postman_collection.json** están los mismos requests en **03_gin_middleware/GET_200** para el caso feliz y **03_gin_middleware/POST_404** para el caso donde no existe el endpoint por si queremos probarlos allí.  
