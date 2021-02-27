package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
	return
}
