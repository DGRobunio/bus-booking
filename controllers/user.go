package controllers

import (
	"bus-booking/models"
	"bus-booking/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NowUser(c *gin.Context) {
	session, _ := c.Cookie("session")
	user := models.User{}
	err := models.NowUser(&user, &session)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"user": gin.H{
			"userId":  user.UserID,
			"account": user.Account,
			"balance": user.Balance,
			"isAdmin": user.IsAdmin,
		},
	})
}

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	user := models.User{Account: account, Password: password}
	session, err := models.Login(&user)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.SetCookie("session", session, 3600, "/", util.Domain, false, false)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"session": session,
	})
}

func Logout(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	models.Logout(&session)
	c.SetCookie("session", session, -1, "/", util.Domain, false, false)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func SignUp(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	user := models.User{Account: account, Password: password}
	err := models.SignUp(&user)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func UpdateUser(c *gin.Context) {
	session, err := c.Cookie("session")
	if err != nil {
		util.Unauthorized(c)
		return
	}
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	user := models.User{Password: oldPassword}
	err = models.UpdateUser(&user, &session, &newPassword)
	if err != nil {
		util.BadRequest(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}
