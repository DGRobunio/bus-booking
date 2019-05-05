package main

import (
	"bus-booking/controllers"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.Default()
	router.GET("/", controllers.NowUser)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.GET("/user", controllers.NowUser)
	router.POST("/user", controllers.SignUp)
	router.PUT("/user", controllers.UpdateUser)
	return router
}
