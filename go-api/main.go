package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"phpguru.net/go-api/db"
	"phpguru.net/go-api/models"
)

func main() {
	db.GetDB()
	server := gin.Default()

	server.GET("/ping", pingHandler)

	server.GET("/events", getEventsHandler)

	server.GET("/events/:id", getEventHandler)

	server.POST("/events", postEventsHandler)

	server.Run(":8080") // localhost:8080
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func getEventsHandler(c *gin.Context) {
	var events, err = models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, events)
}

func postEventsHandler(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	event.UserID = 1
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, event)
}

func getEventHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, event)
}
