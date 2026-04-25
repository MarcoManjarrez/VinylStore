package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var users []User
var signKey = []byte("PleaseGod12345Ilikh")

func CreateAccount(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if json.Username == "" || json.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": "Password or username empty."})
		return
	} else if len(json.Password) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": "Password must be at least 5 characters."})
		return
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error hashing password"})
			return
		} else {
			json.Password = string(hash)
			users = append(users, json)
			c.JSON(http.StatusAccepted, gin.H{"Message": "User registered succesfully"})
		}
	}
}

func Login(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if json.Username == "" || json.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": "Password or username empty."})
		return
	} else {
		var foundUser *User
		for i := range users {
			if users[i].Username == json.Username {
				foundUser = &users[i]
				break
			}
		}

		if foundUser == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "User not found"})
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(json.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error: ": "Password incorrect."})
			return
		} else {
			claims := jwt.MapClaims{}
			claims["user"] = foundUser.Username
			claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
			generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := generatedToken.SignedString(signKey)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
				return
			} else {
				c.JSON(http.StatusAccepted, gin.H{"Message:": "Succesfull login", "Token": tokenString})
			}
		}
	}
}
