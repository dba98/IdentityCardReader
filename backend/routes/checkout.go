package routes

import (
	"IdentityCardReader/backend/controllers"

	// Other imports
	"github.com/gin-gonic/gin"
)

// @Summary Recupera todos os usuários
// @Description Devolve todos os usuários presentes no sistema
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Router /checkout [get]
// @Success 200 {array} model.Users
// @Failure 404 "Not found"
func GetAllUsers(c *gin.Context) {
	controllers.GetAllUsers(c)
}

func DeleteUser(c *gin.Context) {
	controllers.DeleteUser(c)
}
