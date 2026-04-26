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

	//Formats into correct format
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, gin.H{ //Returns the status
		"message": fmt.Sprintf("Hi %s, the DPIP System is Up and Running", username),
		"time":    currentTime,
	})
}
