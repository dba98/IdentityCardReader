package main

import (

	// Other imports

	"PAWPAMFinal/backend/routes"

	"github.com/JsBraz/ProjetoAppWeb/backend/model"
	"github.com/JsBraz/ProjetoAppWeb/backend/services"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var identityKey = "id"

func init() {
	services.OpenDatabase()
	services.Db.AutoMigrate(&model.Location{})
	services.Db.AutoMigrate(&model.Users{})

	defer services.Db.Close()
}

func main() {

	services.FormatSwagger()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// NO AUTH
	//router.GET("/echo/:echo", routes.EchoRepeat)

	// AUTH
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	checkout := router.Group("/api")
	checkout.Use(services.AuthorizationRequired())
	{
		checkout.POST("/addIdentityCardInfo", routes.AddIdentityCardInfo)
		checkout.DELETE("/DeleteIdentityCardInfo/:id", routes.DeleteIdentityCardInfo)
		checkout.GET("/GetIdentityCardInfo/:id", routes.GetIdentityCardInfo)
		checkout.GET("/GetAllIdentityCardInfo", routes.GetAllIdentityCardInfo)
	}

	auth := router.Group("/api/auth")
	{
		// devolve admin=true(Admin) or admin=false(Staff)
		auth.POST("/login", routes.GenerateToken)
		auth.POST("/register", routes.RegisterUser)
		auth.PUT("/refresh_token", services.AuthorizationRequired(), routes.RefreshToken)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
