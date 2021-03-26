# Grupos de rutas en gin

1. En este ejemplo vemos dos grupos de rutas en donde anteponemos **v1** o **v2** en la URL para utilizar distintos handlers.  
2. Para este caso en particular usamos los mismos handlers en los endpoints de **login** y **submit** para los dos grupos pero tenemos un handler distinto en el grupo **v2** del endpoint de **read**.  
3. En la collection de postman **api_rest.postman_collection.json** están los mismos requests en **04_gin_grupos_rutas** para todos los endpoints de ambos grupos **v1** y **v2**.  
