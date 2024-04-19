package routes

import (
	"github.com/DOSuzer/go-jwt-auth/controllers"
	"github.com/DOSuzer/go-jwt-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.POST("/refresh", controllers.Refresh)
	r.GET("/home", middlewares.IsAuthorized(), controllers.Home)
	r.GET("/premium", controllers.Premium)
	r.GET("/me", middlewares.IsAuthorized(), controllers.Me)
	r.PATCH("/me", middlewares.IsAuthorized(), controllers.Update)
	//r.GET("/logout", controllers.Logout)
}
