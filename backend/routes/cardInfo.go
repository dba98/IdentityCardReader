package routes

import (
	"IdentityCardReader/backend/controllers"

	"github.com/gin-gonic/gin"
)

func AddIdentityCardInfo(c *gin.Context) {
	controllers.AddIdentityCardInfo(c)
}

func DeleteIdentityCardInfo(c *gin.Context) {
	controllers.DeleteIdentityCardInfo(c)
}

func GetAllIdentityCardInfo(c *gin.Context) {
	controllers.GetAllIdentityCardInfo(c)
}

func GetIdentityCardInfo(c *gin.Context) {
	controllers.GetIdentityCardInfo(c)
}
