# Usemos paquetes dentro del mismo proyecto

Para utilizar paquetes dentro del mismo proyecto, debemos hacer referencia al path que hay del módulo si usamos go modules (como es nuestro caso), o el path donde nos encontremos después de `$GOPATH/src` si no estamos usando go modules.  

1. Para inicializar nuestro proyecto con go modules, parados en el directorio del proyecto ejecutaremos el comando:  
`go mod init gitlab.grupoesfera.com.ar/capacitacion/CAP-00104-GoModalidadVirtual`  
2. Veremos que dentro del proyecto se crea un archivo go.mod que indica el módulo del proyecto, ¿qué vemos ahí?  
3. En la primera línea de go.mod aparece el módulo que acabamos de inicializar. Cuando queramos importar un paquete que está dentro de nuestro proyecto, en el import debemos escribir la ruta como aparece en go.mod y agregar el directorio donde está nuestro paquete.  
4. Tenemos un ejemplo en el archivo `main.go` de esta carpeta.  
5. Para ejecutarlo, desde el directorio del proyecto ejercutamos:  
`go run 01_conceptos_basicos/02_importar_paquetes_proyecto/main.go`  
6. En este archivo usamos el paquete main para poder ejecutar la función main (¿qué sucede si cambiamos ese nombre de paquete?) e importamos el paquete helloworld que es el que está en el directorio `01_hello_world`.  
7. ¿Qué pasa si descomentamos la línea que llama a "imprimirHelloWorld" (¡prestemos atención a las mayúsculas!).
