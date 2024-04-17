package controllers

import (
	"github.com/DOSuzer/go-jwt-auth/models"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	email, ok := c.Get("email")

	if !ok {
		c.JSON(400, gin.H{"error": "unauthorized"})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", email).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	c.JSON(200, gin.H{
		"id":    existingUser.ID,
		"name":  existingUser.Name,
		"email": existingUser.Email,
		"role":  existingUser.Role})
}

func Update(c *gin.Context) {
	email, ok := c.Get("email")

	if !ok {
		c.JSON(400, gin.H{"error": "unauthorized"})
		return
	}

	var input models.UserNameChange
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", email).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	existingUser.Name = input.Name
	models.DB.Save(&existingUser)

	c.JSON(200, gin.H{
		"id":    existingUser.ID,
		"name":  existingUser.Name,
		"email": existingUser.Email,
		"role":  existingUser.Role})
}
