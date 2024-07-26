package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PasswordGenRequest struct {
	Length                  int  `json:"length"`
	StartWithLetterOrNumber bool `json:"startWithLetterOrNumber"`
}

func main() {
	router := gin.Default()
	router.POST("/generate", generatePassword)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Printf("Error starting router: %s\n", err)
		return
	}
}

func generatePassword(c *gin.Context) {
	var request PasswordGenRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, "")
}
