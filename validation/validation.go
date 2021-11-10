package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pclokcer/dto"
)

func ValidateLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		var loginDto dto.LoginDTO
		if err := c.BindJSON(&loginDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Parametrelerde validasyon yapıldı
		if err := validator.New().Struct(&loginDto); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set("login", loginDto)

		c.Next()
	}
}

func ValidateRegister() gin.HandlerFunc {
	return func(c *gin.Context) {

		var register dto.Register
		if err := c.BindJSON(&register); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Parametrelerde validasyon yapıldı
		if err := validator.New().Struct(&register); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set("register", register)

		c.Next()
	}
}
