package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(logger)

	r.GET("/", funcionQueHaceGet)

	r.Run(":8080")
}

func logger(c *gin.Context) {
	log.Println("Request!")
}

func funcionQueHaceGet(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
