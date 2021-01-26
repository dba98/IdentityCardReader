package controllers

import (
	"IdentityCardReader/backend/model"
	"IdentityCardReader/backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdentityCardInfo(c *gin.Context) {

	var identityCard model.IdentityCard
	var nif string

	if err := c.ShouldBindJSON(&identityCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
	}
	//fmt.Println(c.Keys("nif"))
	//	fmt.Println(c.Keys)
	fmt.Println(nif+ "AQUI AQUI AQUI AQUI")
	services.Db.Find(&identityCard, "nif = ?", identityCard.Nif)

	fmt.Println(identityCard)
	if identityCard.Nif == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid IdentityCard!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "id": identityCard.ID, "nif": identityCard.Nif, "frontData": identityCard.Image1, "backData": identityCard.Image2})
	return
}

func AddIdentityCardInfo(c *gin.Context) {
	var identityCard model.IdentityCard

	if err := c.ShouldBindJSON(&identityCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	identityCard.Image1 = services.CardImageDecrypt(identityCard.Image1)
	identityCard.Image2 = services.CardImageDecrypt(identityCard.Image2)

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

	nif := c.Param("id")
	services.Db.First(&identityCard, nif)

	if identityCard.Nif == "0"  {
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