package controllers

import (
	"IdentityCardReader/backend/model"
	"IdentityCardReader/backend/services"
	"fmt"

	// Local imports
	"regexp"

	// Other imports
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var creds model.User
	var usr model.User

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}
	services.OpenDatabase()
	services.Db.Find(&usr, "username = ? and password = ?", creds.Username, creds.Password)
	if usr.Username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid User!"})
		return
	}

	token, expirationTime, isAdmin := services.GenerateTokenJWT(usr)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Access denied!"})
		return
	}
	isAdmin = true
	defer services.Db.Close()
	fmt.Println(usr.Username + " teste")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "id": usr.ID, "username": usr.Username, "token": token, "expirationTime": expirationTime, "isAdmin": isAdmin, "nif": usr.Nif})
}

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func RegisterHandler(c *gin.Context) {
	var creds model.User

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad request!"})
		return
	}

	if !isEmailValid(creds.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid email"})
		return
	}

	services.OpenDatabase()
	value := services.Db.Save(&creds)
	if value.RowsAffected == 0 {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Username already taken!"})
		defer services.Db.Close()
		return
	}
	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success!", "ID": creds.ID, "user": creds})
}

func RefreshHandler(c *gin.Context) {

	user := model.User{
		Username: c.GetHeader("username"),
	}

	token, expirationTime, isAdmin := services.GenerateTokenJWT(user)

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Acesso não autorizado"})
		return
	}

	defer services.Db.Close()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusNoContent, "message": "Token atualizado com sucesso!", "token": token, "expirationTime": expirationTime, "isAdmin": isAdmin})
}
