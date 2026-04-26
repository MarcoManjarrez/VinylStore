package main

import (
	"log"

	"github.com/MarcoManjarrez/VinylStore.git/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Public endpoints
	r.POST("/createAccount", handlers.CreateAccount) // Utility to register users
	r.GET("/login", handlers.Login)                  // Login using Basic Auth or JSON
	r.POST("/login", handlers.Login)                 // Support both GET and POST for login flexibility

	// Protected endpoints
	protected := r.Group("/")
	protected.Use(handlers.CheckAuthorization()) // Middleware to protect endpoints
	{
		protected.GET("/logout", handlers.Logout)           // Revoke token
		protected.POST("/logout", handlers.Logout)          // Support both GET and POST
		protected.GET("/albums", handlers.GetAlbums)        // Get all albums
		protected.GET("/albums/:id", handlers.GetAlbumByID) // Get album by specific ID
		protected.POST("/post-album", handlers.PostAlbum)   // Add a new album (per curl example)
		protected.POST("/createAlbum", handlers.PostAlbum)  // Add a new album (per list example)
		protected.GET("/status", handlers.GetSystemStatus)  // Get system status and user
	}

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
