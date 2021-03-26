# Servidor HTTP

1. Este es un ejemplo de cómo, únicamente usando el paquete `http` de golang, podemos levantar un servidor http de manera simple.  
2. Aquí exponemos dos endpoints, uno que nos saluda y otro que muestra las cabeceras del request http que llegó.  
3. El concepto de **handler** se utiliza mucho para este paquete, que son los tipos que implementan la interfaz `http.Handler`.  
4. Una manera común de hacer esto es:
   1. Definir funciones que tengan como entrada `(http.ResponseWriter, *http.Request)` y no tengan salida.
   2. La idea de estas funciones es que manejen el request que llega como segundo parámetro en la entrada y escriban la respuesta utilizando el primer parámetro de entrada.
   3. Pasar esa función por parámetro a la función `http.HandleFunc` que recibe un patrón de la URL y la función que nombrábamos para registrar un handler nuevo.  
5. Cuando ejecutamos este main, quedará levantado un server al que podremos hacerle requests:
   1. Un request posible es un GET a localhost:8090/hello que imprime un "hello" y desde la terminal podemos hacer con `curl localhost:8090/hello`, ir a esa url desde un browser o enviar el request por Postman.  
   2. Otro request posible es un GET a localhost:8090/headers que imprime los headers del request y desde la terminal podemos hacer con `curl localhost:8090/headers`, ir a la url desde un browser o enviar el request por Postman (los headers serán distintos en cada caso).  
6. En la collection de postman **api_rest.postman_collection.json** están los mismos requests en **02_servidor_http/GET_hello** y **02_servidor_http/GET_headers** por si queremos probarlos allí.  
