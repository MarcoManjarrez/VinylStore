package main

import (
	"log"

	"github.com/MarcoManjarrez/VinylStore.git/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Public endpoints
	r.POST("/createAccount", handlers.CreateAccount) //Registers users
	r.GET("/login", handlers.Login)                  //Login using Basic Auth or JSON
	r.POST("/login", handlers.Login)                 //More flexible call to login with post

	//Protected endpoints
	protected := r.Group("/")
	protected.Use(handlers.CheckAuthorization()) //Middleware to protect endpoints behind auth tokens
	{
		protected.GET("/logout", handlers.Logout)           //Revoke token
		protected.POST("/logout", handlers.Logout)          //Sam flexibility for logout
		protected.GET("/albums", handlers.GetAlbums)        //Get all albums
		protected.GET("/albums/:id", handlers.GetAlbumByID) //Get album with specific ID
		protected.POST("/post-album", handlers.PostAlbum)   //Add a new album (curl)
		protected.POST("/createAlbum", handlers.PostAlbum)  //Add a new album (list)
		protected.GET("/status", handlers.GetSystemStatus)  //Get system status and user
	}

	// Starts the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
