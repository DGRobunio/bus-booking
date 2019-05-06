package controllers

import (
	"bus-booking/models"
	"bus-booking/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AllFavorites(c *gin.Context) {
	session, _ := c.Cookie("session")
	buses := make([]models.Bus, 0)
	err := models.AllFavorites(&session, &buses)
	if err != nil {
		util.BadRequest(c)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"bus":  buses,
	})
}

func Favorite(c *gin.Context) {
	session, _ := c.Cookie("session")
	busID := c.PostForm("busID")
	err := models.Favorite(&session, &busID)
	if err != nil {
		util.BadRequest(c)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func Unfavorite(c *gin.Context) {
	session, _ := c.Cookie("session")
	busID := c.Query("busID")
	log.Print(session)
	log.Print(busID)
	err := models.Unfavorite(&session, &busID)
	if err != nil {
		util.BadRequest(c)
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}
