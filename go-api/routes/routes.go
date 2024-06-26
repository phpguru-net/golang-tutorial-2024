package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEventsHandler)
	server.GET("/events/:id", getEventHandler)
	server.POST("/events", postEventsHandler)
	server.PUT("/events/:id", updateEventHandler)
	server.DELETE("/events/:id", deleteEventHandler)
}
