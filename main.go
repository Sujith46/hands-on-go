package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Text string `json:"message"`
}

func main() {
	route := gin.Default()

	route.GET("/", helloHandler)
	route.POST("/post", postHandler)
	route.GET("/get", getHandler)
	err := route.Run(":8080")
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func helloHandler(c *gin.Context) {
	message := Message{Text: "Hello, World!"}
	c.JSON(http.StatusOK, message)
}

func postHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	var requestBody Message
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	response := Message{Text: "Received: " + requestBody.Text}
	c.JSON(http.StatusOK, response)
}

func getHandler(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	message := Message{Text: "This is a GET request!"}
	c.JSON(http.StatusOK, message)
}
