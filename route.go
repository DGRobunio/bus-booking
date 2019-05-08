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
	router.GET("/bus", controllers.AllBuses)
	router.GET("/bus/:busID", controllers.OneBus)
	router.GET("/order", controllers.AllOrders)
	router.GET("/order/:orderID", controllers.OneOrder)
	router.POST("/order", controllers.Book)
	router.GET("/favorite", controllers.AllFavorites)
	router.POST("/favorite", controllers.Favorite)
	router.DELETE("/favorite", controllers.Unfavorite)
	return router
}
