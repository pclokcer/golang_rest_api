package main

import "github.com/gin-gonic/gin"
import "gorm.io/gorm"
import "github.com/pclokcer/config"
import "github.com/pclokcer/controller"
import "github.com/pclokcer/middleware"
import "github.com/pclokcer/service"

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	jwtService     service.JWTService        = service.NewJWTService()
	authController controller.AuthController = controller.NewAuthController()
)

func main()  {
	r:= gin.Default()

	authRoutes := r.Group("/api/auth", middleware.AuthorizeJWT(jwtService))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.GET("/test", authController.Register)
	}

	r.Run(":3000")
}