package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil { //Maps json to the User var
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.Username == "" || json.Password == "" { //Checks if both Username and Password are passed. BadRequest if not
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password or username cannot be empty."})
		return
	}
	if len(json.Password) < 5 { //Checks if password is 5 characters in length or more. BadRequest if not
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 5 characters."})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(json.Password), 10) //Attempts to hash pasword ebfore storing. InternalServerError if not
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	mu.Lock() //Protects resources to first change the string password to the generated hash and then stores the new user in the slice
	json.Password = string(hash)
	users = append(users, json)
	mu.Unlock()

	c.JSON(http.StatusAccepted, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	username, password, hasAuth := c.Request.BasicAuth() //Gets usrname, password and hasAuth from request

	if !hasAuth { //If no hasAuth exists, we have a backup json request attempt
		var json User
		if err := c.ShouldBindJSON(&json); err == nil { //Binds the json to the variable to see if we can rescue the request
			username = json.Username
			password = json.Password
		}
	}

	if username == "" || password == "" { //Checks if username and password exists. BR if not
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password or username empty."})
		return
	}

	var foundUser *User //Check to find the username
	for i := range users {
		if users[i].Username == username {
			foundUser = &users[i]
			break
		}
	}

	if foundUser == nil { //Check to see if the username does exist or not. Status unauthorized if not
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)) //Compares the found user's hashed password with the string password from request to validate
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password incorrect."})
		return
	}

	claims := jwt.MapClaims{ //Maps a jwt claim to generate an access token
		"user": foundUser.Username,                   //Keep username the same
		"exp":  time.Now().Add(time.Hour * 2).Unix(), //Expiration 2 hours from creation
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //Gets token and turns it to string
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

// Logs user out
func Logout(c *gin.Context) {
	tokenString := c.GetString("raw_token") //Gets token and username
	username := c.GetString("username")

	mu.Lock()
	revokedTokens[tokenString] = true //Turns token into a revokedToken
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Bye %s, your token has been revoked", username),
	})
}
