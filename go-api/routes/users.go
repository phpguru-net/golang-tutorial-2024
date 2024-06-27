package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"phpguru.net/go-api/models"
)

func postRegisterUserHandler(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can not parse user data",
		})
		return
	}
	err = user.CreateUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": user.Id,
	})
}

func postLoginHandler(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can not parse json",
		})
		return
	}
	isValidPassword, err := user.ValidateCredentials()
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server error",
		})
		return
	}
	if isValidPassword {
		accessToken, err := user.GenerateAccessTokens()
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"accessToken": accessToken,
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Invalid email or password!",
	})
}
