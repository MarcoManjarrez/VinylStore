package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// CheckAuthorization validates the JWT token
func CheckAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an authorization token"})
			c.Abort()
			return
		}

		//Extracts token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		//Fallback in case the user didn't include the "Bearer " prefix
		if tokenString == authHeader {
			tokenString = strings.TrimSpace(authHeader)
		}

		//Check if token has been revoked
		mu.Lock()
		isRevoked := revokedTokens[tokenString]
		mu.Unlock()

		if isRevoked {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been revoked. Please login again."})
			c.Abort()
			return
		}

		//Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return signKey, nil
		})

		//Validate token and extract claims
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token detected"})
			c.Abort()
			return
		}

		//Set variables into context for endpoints to use
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["user"])
			c.Set("raw_token", tokenString)
		}

		c.Next()
	}
}
