package routes

import (
	"IdentityCardReader/backend/controllers"

	"github.com/gin-gonic/gin"
)

/**
 * Cada rota vai buscar a respetiva função à package dos controladores
**/

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
