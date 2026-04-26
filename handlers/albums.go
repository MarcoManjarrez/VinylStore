package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums) //Simply returns all albums in albums slice
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id") //Gets the id in the params

	for _, album := range albums { //Checks all albums to see if one has a corresponding one
		if album.ID == id {
			c.JSON(http.StatusOK, []Album{album}) //Success status
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
}

// Posts new album to the registry
func PostAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //If binding json fields fails returns a bad request error
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, album := range albums { //Accesses the album registry to check if the new album's id is unique
		if album.ID == newAlbum.ID {
			c.JSON(http.StatusConflict, gin.H{"error": "An album with this ID already exists"}) //Returns a conflict error if not
			return
		}
	}

	albums = append(albums, newAlbum)    //Adds new album
	c.JSON(http.StatusCreated, newAlbum) //Success
}
