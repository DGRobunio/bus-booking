package controllers

import (
	"bus-booking/models"
	"bus-booking/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllOrders(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	orders := make([]models.Order, 0)
	err = models.AllOrders(&session, &orders)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"order": orders,
	})
}

func OneOrder(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	order := models.Order{OrderID: c.Param("orderID")}
	err = models.OneOrder(&session, &order)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"order": order,
	})
}
