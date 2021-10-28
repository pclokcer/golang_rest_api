package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclokcer/dto"
	"github.com/pclokcer/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	connection *mongo.Database
}

func NewAuthController(mongo *mongo.Database) AuthController {
	return &authController{
		connection: mongo,
	}
}

func (auth *authController) Login(c *gin.Context) {

	var loginDto dto.LoginDTO

	err := c.BindJSON(&loginDto)

	if err != nil {
		panic(err)
	}

	var loginWithsDto dto.LoginWithsDTO

	err = auth.connection.Collection("login_withs").FindOne(context.TODO(), bson.D{{"email", loginDto.Email}}).Decode(&loginWithsDto)

	if loginWithsDto.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Kullanıcı Adı veya Şifre Yanlış",
		})
		return
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}

	if nil != bcrypt.CompareHashAndPassword([]byte(loginWithsDto.Password), []byte(loginDto.Password)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Kullanıcı Adı veya Şifre Yanlış",
		})
		return
	}

	token := service.NewJWTService().GenarateToken("erdem")

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"email": loginDto.Email,
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registered",
	})
}
