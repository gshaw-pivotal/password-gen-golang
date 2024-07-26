package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type PasswordGenRequest struct {
	Length                  int  `json:"length"`
	StartWithLetterOrNumber bool `json:"startWithLetterOrNumber"`
}

var LowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"

var UpperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var specialCharacters = "!@#$%^&*-_?"

var specialCharactersLength = len(specialCharacters)

func main() {
	router := gin.Default()
	router.POST("/generate", generatePasswordRequestHandler)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Printf("Error starting router: %s\n", err)
		return
	}
}

func generatePasswordRequestHandler(c *gin.Context) {
	var request PasswordGenRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var password = generatePassword(request.Length, request.StartWithLetterOrNumber)

	c.IndentedJSON(http.StatusOK, password)
}

func generatePassword(length int, startWithLetterOrNumber bool) string {

	var password strings.Builder

	var charType int

	for i := 0; i < length; i++ {
		if startWithLetterOrNumber {
			charType = rand.Intn(3)
		} else {
			charType = rand.Intn(4)
		}

		switch charType {
		case 0:
			// Lowercase letter
			password.WriteString(string(getChar(LowerCaseLetters, rand.Intn(26))))
		case 1:
			// Uppercase letter
			password.WriteString(string(getChar(UpperCaseLetters, rand.Intn(26))))
		case 2:
			// Number
			password.WriteString(strconv.Itoa(rand.Intn(10)))
		case 3:
			// Special char
			password.WriteString(string(getChar(specialCharacters, rand.Intn(specialCharactersLength))))
		}
	}

	return password.String()
}

func getChar(s string, index int) rune {
	return []rune(s)[index]
}
