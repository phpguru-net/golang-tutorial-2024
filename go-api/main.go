package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"phpguru.net/go-api/db"
	"phpguru.net/go-api/routes"
)

func main() {
	db.GetDB()
	server := gin.Default()

	server.GET("/ping", pingHandler)

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
