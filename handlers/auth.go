package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CreateAccount allows a new user to register
func CreateAccount(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Username == "" || json.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password or username cannot be empty."})
		return
	}
	if len(json.Password) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 5 characters."})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	mu.Lock()
	json.Password = string(hash)
	users = append(users, json)
	mu.Unlock()

	c.JSON(http.StatusAccepted, gin.H{"message": "User registered successfully"})
}

// Login authenticates the user and provides a token
func Login(c *gin.Context) {
	// The instructions show curl -u username:password, which implies Basic Auth.
	username, password, hasAuth := c.Request.BasicAuth()

	// Fallback to JSON payload if Basic Auth is not provided
	if !hasAuth {
		var json User
		if err := c.ShouldBindJSON(&json); err == nil {
			username = json.Username
			password = json.Password
		}
	}

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password or username empty."})
		return
	}

	var foundUser *User
	for i := range users {
		if users[i].Username == username {
			foundUser = &users[i]
			break
		}
	}

	if foundUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password incorrect."})
		return
	}

	// Generate Token
	claims := jwt.MapClaims{
		"user": foundUser.Username,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hi %s, welcome to the Store System", foundUser.Username),
		"token":   tokenString,
	})
}

// Logout revokes the token so it cannot be used again
func Logout(c *gin.Context) {
	tokenString := c.GetString("raw_token")
	username := c.GetString("username")

	mu.Lock()
	revokedTokens[tokenString] = true
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Bye %s, your token has been revoked", username),
	})
}
