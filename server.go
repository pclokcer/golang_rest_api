package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pclokcer/config"
	"github.com/pclokcer/controller"
	"github.com/pclokcer/middleware"
	"github.com/pclokcer/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	jwtService     service.JWTService        = service.NewJWTService()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	req := gin.Default()

	authRoutes := req.Group("/api")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	generalRequest := req.Group("/api/", middleware.AuthorizeJWT(jwtService))
	{
		generalRequest.GET("/users", authController.Login)
	}

	req.Run(":3000")
}
