package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pclokcer/dto"
	"github.com/pclokcer/entity"
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

	var loginDto dto.LoginDTO

	err := c.BindJSON(&loginDto)

	if err != nil {
		panic(err)
	}

	// Parametrelerde validasyon yapıldı
	if err := validator.New().Struct(&loginDto); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Kullanıcı Bizim tarafımızda var mı diye kontrol edilyor
	var loginWithsDto dto.LoginWithsDTO
	err = auth.connection.Collection("login_withs").FindOne(context.TODO(), bson.D{{"email", loginDto.Email}}).Decode(&loginWithsDto)

	if loginWithsDto.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{MessageID: "LoginError"}),
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

	// Kullanıcı Şifre comapre ediliyor
	if nil != bcrypt.CompareHashAndPassword([]byte(loginWithsDto.Password), []byte(loginDto.Password)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{MessageID: "LoginError"}),
		})
		return
	}

	// Token üretiliyor
	token := service.NewJWTService().GenarateToken(loginWithsDto.ID)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"email": loginDto.Email,
	})
}

func (auth *authController) Register(c *gin.Context) {

	// dil belirleniyor
	localizer := libs.GetLocalizer(c)

	var login_with entity.LoginWith
	err := c.BindJSON(&login_with)

	if err != nil {
		panic(err)
	}

	// Parametrelerde validasyon yapıldı
	if err := validator.New().Struct(&login_with); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Bu Kullanıcı Zaten Üye mi Diye bakılıyor
	var loginWithsDto dto.LoginWithsDTO
	err = auth.connection.Collection("login_withs").FindOne(context.TODO(), bson.D{{"email", login_with.Email}}).Decode(&loginWithsDto)

	if loginWithsDto.Email != "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{MessageID: "ExistUser"}),
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(login_with.Password), 10)

	login_with.Password = string(hashPassword)

	// Kullanıcı Kaydı Yapılıyor
	res, err := auth.connection.Collection("login_withs").InsertOne(context.Background(), login_with)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	c.JSON(http.StatusOK, gin.H{"success": localizer.(*i18n.Localizer).MustLocalize(&i18n.LocalizeConfig{MessageID: "NewPerson"})})

}
