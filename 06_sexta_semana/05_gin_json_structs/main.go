package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Usuario representa un usuario con email y contraseña
type Usuario struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Correo      string `json:"email" binding:"required"`
	Contrasenia string `json:"password" binding:"required"`
}

var usuarios map[int]Usuario

func main() {
	usuarios = make(map[int]Usuario)
	usuarios[82] = Usuario{ID: 82, Nombre: "Maria", Correo: "maria@ejemplo.com", Contrasenia: "maria"}
	usuarios[3] = Usuario{ID: 3, Nombre: "Ana", Correo: "ana@ejemplo.com", Contrasenia: "ana"}
	usuarios[24] = Usuario{ID: 24, Nombre: "Mario", Correo: "mario@ejemplo.com", Contrasenia: "mario"}

	router := gin.Default()

	router.GET("/usuarios", obtenerUsuarios)
	router.GET("/usuario/:id", obtenerUsuario)
	router.POST("/usuario", agregarUsuario)
	router.DELETE("/usuario/:id", eliminarUsuario)

	router.Run(":8080")
}

func obtenerUsuarios(c *gin.Context) {

	usuariosADevolver := make([]Usuario, 0)
	nombre := c.Query("nombre")

	for _, usuario := range usuarios {
		if nombre == "" || strings.ToLower(usuario.Nombre) == strings.ToLower(nombre) {
			usuariosADevolver = append(usuariosADevolver, usuario)
		}
	}

	c.JSON(http.StatusOK, usuariosADevolver)
}

func obtenerUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuarioADevolver, existe := usuarios[id]

	if !existe {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no existe el usuario con ese id"})
		return
	}

	c.JSON(http.StatusOK, usuarioADevolver)
}

func agregarUsuario(c *gin.Context) {

	var nuevoUsuario Usuario
	nuevoID := obtenerProximoIDDeUsuario()

	err := c.ShouldBindJSON(&nuevoUsuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nuevoUsuario.ID = nuevoID

	usuarios[nuevoID] = nuevoUsuario

	c.String(http.StatusCreated, strconv.Itoa(nuevoID))
}

func obtenerProximoIDDeUsuario() int {
	ultimoID := 0

	for id := range usuarios {
		if id > ultimoID {
			ultimoID = id
		}
	}

	return ultimoID + 1
}

func eliminarUsuario(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delete(usuarios, id)

	c.Status(http.StatusNoContent)
}
