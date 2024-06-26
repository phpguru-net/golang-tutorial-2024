package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"phpguru.net/go-api/models"
)

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
		c.JSON(http.StatusBadRequest, "Could not parse data")
		return
	}
	event.UserID = 1
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
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

func updateEventHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	_, err = models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data",
		})
		return
	}
	event.ID = id

	err = event.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, event)
}

func deleteEventHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Can not parse id",
		})
		return
	}
	err = models.DeleteEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Some went wrong " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
