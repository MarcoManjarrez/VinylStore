package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Request does not contain an authorization token"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { //Checa si el signing es hmac. Si no es, regresa que token es nil y un error
				return nil, fmt.Errorf("unexpected signing erro")
			}
			return signKey, nil
		})

		if err != nil || !token.Valid { //Checa si hubo error o si el token no es valido
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token detected"})
			c.Abort()
			return
		}

		c.Next()
	}
}
