package routes

import (
	"github.com/gin-gonic/gin"
	"phpguru.net/go-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	privateGroups := server.Group("/")
	privateGroups.Use(middlewares.Authorization)

	privateGroups.GET("/events", getEventsHandler)
	privateGroups.GET("/events/:id", middlewares.Authorization, getEventHandler)
	privateGroups.POST("/events", middlewares.Authorization, postEventsHandler)
	privateGroups.PUT("/events/:id", middlewares.Authorization, updateEventHandler)
	privateGroups.DELETE("/events/:id", middlewares.Authorization, deleteEventHandler)

	server.POST("/register", postRegisterUserHandler)
	server.POST("/login", postLoginHandler)
}
