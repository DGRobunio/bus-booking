package controllers

import (
	"bus-booking/models"
	"bus-booking/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllBuses(c *gin.Context) {
	buses := make([]models.Bus, 0)
	err := models.AllBuses(&buses)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"bus":  buses,
	})
}

func OneBus(c *gin.Context) {
	bus := models.Bus{BusID: c.Param("busID")}
	err := models.OneBus(&bus)
	if err != nil {
		util.BadRequest(c)
		return
	}
	session, _ := c.Cookie("session")
	if session != "" {
		var user models.User
		err := models.NowUser(&user, &session)
		util.Report(err)
		favorited := models.Favorited(&user.UserID, &bus.BusID)
		c.JSON(http.StatusOK, gin.H{
			"code":      http.StatusOK,
			"bus":       bus,
			"favorited": favorited,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"bus":  bus,
		})
	}
}
