package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetSystemStatus returns system status and logged user details
func GetSystemStatus(c *gin.Context) {
	username := c.GetString("username")

	// Format matches PDF: "2015-03-07 11:06:39"
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hi %s, the DPIP System is Up and Running", username),
		"time":    currentTime,
	})
}
