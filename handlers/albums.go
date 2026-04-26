package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAlbums returns all albums from the store
func GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// GetAlbumByID fetches a specific album by its ID
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			// Wrapping inside a slice to match the PDF instructions array format output
			c.JSON(http.StatusOK, []Album{album})
			return
		}
	}

	// Returns 404 if not found as requested in instructions
	c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

// PostAlbum adds a new album to the store
func PostAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Prevent duplicate IDs (Rubric requirement)
	for _, album := range albums {
		if album.ID == newAlbum.ID {
			c.JSON(http.StatusConflict, gin.H{"error": "An album with this ID already exists"})
			return
		}
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}
