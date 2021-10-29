package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pclokcer/config"
	"github.com/pclokcer/controller"
	"github.com/pclokcer/middleware"
	"github.com/pclokcer/service"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	mongoDB           *mongo.Database              = config.MongoConnection()
	jwtService        service.JWTService           = service.NewJWTService()
	authController    controller.AuthController    = controller.NewAuthController(mongoDB)
	commentController controller.CommentController = controller.NewCommentController(mongoDB)
	uploadController  controller.UploadController  = controller.NewUploadController()
)

func main() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Env file error")
	}

	req := gin.Default()

	authRoutes := req.Group("/api")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	generalRequest := req.Group("/api", middleware.AuthorizeJWT(jwtService))
	{
		generalRequest.POST("/users/:param", controller.GetUsers)
		generalRequest.POST("/get-comments", commentController.All)
		generalRequest.POST("/set-comment", commentController.SetComment)
		generalRequest.POST("/image-upload", uploadController.Upload)
	}

	req.Run(":" + os.Getenv("PORT"))
}
