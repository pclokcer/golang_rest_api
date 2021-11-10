package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pclokcer/dto"
	"github.com/pclokcer/libs"
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

	// dil belirleniyor
	localizer := libs.GetLocalizer(c)

	loginDto, _ := c.Get("login")

	// Kullanıcı Bizim tarafımızda var mı diye kontrol edilyor
	var loginWithsDto dto.LoginWithsDTO
	auth.connection.Collection("login_withs").FindOne(context.TODO(), bson.D{{"email", loginDto.(dto.LoginDTO).Email}}).Decode(&loginWithsDto)

	if loginWithsDto.Email == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{MessageID: "LoginError"}),
		})
		return
	}

	// Kullanıcı Şifre comapre ediliyor
	if nil != bcrypt.CompareHashAndPassword([]byte(loginWithsDto.Password), []byte(loginDto.(dto.LoginDTO).Password)) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{MessageID: "LoginError"}),
		})
		return
	}

	// Token üretiliyor
	token := service.NewJWTService().GenarateToken(loginWithsDto.ID)

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"token": token,
		"email": loginDto.(dto.LoginDTO).Email,
	})
}

func (auth *authController) Register(c *gin.Context) {

	// dil belirleniyor
	localizer := libs.GetLocalizer(c)

	register, _ := c.Get("register")

	// Bu Kullanıcı Zaten Üye mi Diye bakılıyor
	var loginWithsDto dto.LoginWithsDTO
	auth.connection.Collection("login_withs").FindOne(context.TODO(), bson.D{{"email", register.(dto.Register).Email}}).Decode(&loginWithsDto)

	if loginWithsDto.Email != "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{MessageID: "ExistUser"}),
		})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(register.(dto.Register).Password), 10)

	var regiserData dto.LoginWithsDTO
	regiserData.Password = string(hashPassword)
	regiserData.Email = register.(dto.Register).Email

	// Kullanıcı Kaydı Yapılıyor
	auth.connection.Collection("login_withs").InsertOne(context.Background(), regiserData)

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{MessageID: "NewPerson"})})

}
