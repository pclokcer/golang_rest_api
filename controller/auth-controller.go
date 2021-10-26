package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclokcer/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {

	token := service.NewJWTService().GenarateToken("erdem")

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registered",
	})
}
