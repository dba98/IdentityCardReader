package controllers

import (
	"IdentityCardReader/backend/model"
	"IdentityCardReader/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetIdentityCardInfo(c *gin.Context) {

	var identityCard model.IdentityCard
	var creds model.IdentityCard


	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
	}
	services.Db.Find(&identityCard, "identityCardID = ?", creds.ID)

	if identityCard.Nif == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid IdentityCard!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "Username": creds.Nif})
	return
}

func AddIdentityCardInfo(c *gin.Context){
	var identityCard model.IdentityCard

	if err := c.ShouldBindJSON(&identityCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	value := services.Db.Save(&identityCard)
	if value.RowsAffected == 0 {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "nif already taken!"})
		defer services.Db.Close()
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "ID": identityCard.ID, "identityCard": identityCard})


}

func DeleteIdentityCardInfo(c *gin.Context) {
	var identityCard model.IdentityCard

	id := c.Param("id")
	services.Db.First(&identityCard, id)

	if identityCard.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&identityCard)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}



func GetAllIdentityCardInfo(c *gin.Context) {
	var identityCards []model.IdentityCard

	services.Db.Find(&identityCards)

	if len(identityCards) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": identityCards})
}
