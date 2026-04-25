package handlers

import (
	//"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VinylRegistry struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseJson struct {
	Message string          `json:"message"`
	Content []VinylRegistry `json:"content"`
}

var posts []VinylRegistry

func GetVinylRegistry(c *gin.Context) {
	var response responseJson
	response.Content = posts
	c.JSON(http.StatusAccepted, response)
}
