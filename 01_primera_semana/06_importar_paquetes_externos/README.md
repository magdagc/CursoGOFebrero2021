# Veamos cómo manejar las dependencias con go modules

1. Hay un test que hace referencia a una función que todavía no existe, en este test hacemos lo mismo de dos formas distintas.
   1. El primer test hace uso del paquete `testing` propio de golang.
   2. El segundo test hace uso del paquete `assert` que obtenemos del [repositorio de github](https://github.com/stretchr/testify). Este paquete lo vamos a poder usar cuando actualicemos las dependencias con go modules.
2. Borremos el archivo `go.mod`.
3. Intentemos ejecutar el test con el siguiente comando desde el directorio del proyecto y veamos qué sucede:
`go test github.com/magdagc/CursoGOFebrero2021/01_primera_semana/06_importar_paquetes_externos/`
4. Para inicializar go modules y decirle cuál es el módulo de nuestro proyecto **(porque borramos el archivo go.mod en el paso 2)** ejecutamos desde el directorio del proyecto:  
`go mod init github.com/magdagc/CursoGOFebrero2021`
5. Para bajar las dependencias que yo ya haya importado y limpiar si hubiera algunas que no se usen:  
`go mod tidy`
6. Como en mi import de `porcentaje_test.go` ya tengo la referencia al módulo de testify (y el paquete assert dentro del mismo), go modules detecta y descarga esa dependencia. Podemos ver datos de esa dependencia en `go.mod` y en `go.sum` que es otro archivo que se generó, miremos qué hay ahí.
7. ¿Cómo sabe go modules que ese import corresponde a lo que quiero bajar? Deberíamos usar el mismo path que tiene el módulo en el archivo `go.mod` del [repositorio de github](https://github.com/stretchr/testify).
8. Volvamos a intentar ejecutar las pruebas:
`go test github.com/magdagc/CursoGOFebrero2021/01_conceptos_basicos/06_importar_paquetes_externos/`
9. Ahora no tenemos más problema con la dependencia del paquete de assert, pero tenemos que hacer que pase el test.  
