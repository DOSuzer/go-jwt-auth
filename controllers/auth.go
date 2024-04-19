package controllers

import (
	"os"
	"strings"
	"time"

	"github.com/DOSuzer/go-jwt-auth/models"

	"github.com/DOSuzer/go-jwt-auth/utils"

	"github.com/gin-gonic/gin"
)

var secret = os.Getenv("SECRET_KEY")
var refreshSecret = os.Getenv("REFRESH_SECRET_KEY")

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	t, rt, err := utils.CreateTokensForUser(existingUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate tokens"})
		return
	}

	//c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"token": t, "refresh_token": rt})
}

func Refresh(c *gin.Context) {
	var input models.RefreshToken
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}

	isBlacklisted, err := utils.IsBlacklisted(input.RefreshToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not check if token is blacklisted"})
		return
	}
	if isBlacklisted {
		c.JSON(401, gin.H{"error": "token is blacklisted"})
		return
	}

	claims, err := utils.ParseRefreshToken(input.RefreshToken, refreshSecret)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	err = utils.AddToBlacklist(input.RefreshToken, 7*24*time.Hour)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not add token to blacklist"})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", claims.Subject).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	t, rt, err := utils.CreateTokensForUser(existingUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate tokens"})
		return
	}

	//c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"token": t, "refresh_token": rt})
}

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	models.DB.Create(&user)

	c.JSON(200, gin.H{"success": "user created"})
}

func Home(c *gin.Context) {

	role, ok := c.Get("role")
	if !ok {
		c.JSON(500, gin.H{"error": "could not get role"})
	}
	c.JSON(200, gin.H{"success": "home page", "role": role})
}

func Premium(c *gin.Context) {

	authHeader := c.Request.Header.Get("Authorization")
	t := strings.Split(authHeader, " ")

	if len(t) != 2 {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	token := t[1]

	if token == "" {
		c.JSON(401, gin.H{"error": "invalid token"})
		return
	}

	claims, err := utils.ParseToken(token, secret)

	if err != nil {
		c.JSON(401, gin.H{"error": "could not parse token", "message": err.Error()})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "premium page", "role": claims.Role})
}
