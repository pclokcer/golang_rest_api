package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pclokcer/helper"
	"github.com/pclokcer/service"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := helper.BuildErrorResponse("JWT: invalid Signature", "No Token Found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {

			claims := token.Claims.(jwt.MapClaims)
			log.Println("claim[user_id]: ", claims["user_id"])
			log.Println("claim[issuer]: ", claims["issuer"])

		} else {

			log.Println(err)
			response := helper.BuildErrorResponse("JWT: invalid Signature", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)

		}
	}
}
