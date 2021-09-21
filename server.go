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
	r := gin.Default()

	authRoutes := r.Group("/api/auth", middleware.AuthorizeJWT(jwtService))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.GET("/test", authController.Register)
	}

	r.Run(":3000")
}
