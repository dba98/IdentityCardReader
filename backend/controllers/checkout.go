package controllers

import (
	// Local imports

	"IdentityCardReader/backend/model"
	"IdentityCardReader/backend/services"

	// Other imports
	"net/http"

	"github.com/gin-gonic/gin"
)

func Echo(c *gin.Context) {
	echo := c.Param("echo")

	c.JSON(http.StatusOK, gin.H{
		"echo": echo,
	})
}

func DeleteUser(c *gin.Context) {
	var user model.User

	id := c.Param("id")
	services.Db.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}

func GetAllUsers(c *gin.Context) {

	var users []model.User

	services.Db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}
