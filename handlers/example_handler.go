package handlers

import (
	"fmt"
	"net/http"

	"mentalartsapi/dto"

	"github.com/gin-gonic/gin"
)

func HandlePing(c *gin.Context) {
	res := dto.Response{Msg: "pong"}
	c.JSON(http.StatusOK, res)
}

func HandleHello(c *gin.Context) {
	name := c.Query("name")

	var msg string
	if name != "" {
		msg = fmt.Sprintf("Welcome, %s", name)
	} else {
		msg = "Welcome, user"
	}

	c.String(http.StatusOK, msg)
}

func HandleHelloWithPayload(c *gin.Context) {
	// Binding (Get Payload from request)
	var user dto.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	// Validaton (Validate payload)
	if user.Name == "" {
		c.String(http.StatusBadRequest, "empty name is not accepted")
		return
	}

	msg := fmt.Sprintf("Welcome, %s %s", user.Name, user.Surname)
	c.String(http.StatusOK, msg)

}
