package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Grupo v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}
	// Grupo v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpointNuevaVersion)
	}
	router.Run(":8080")
}

func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "loginEndpoint")
}

func submitEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "submitEndpoint")
}

func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "readEndpoint")
}

func readEndpointNuevaVersion(c *gin.Context) {
	c.String(http.StatusOK, "readEndpointNuevaVersion")
}
