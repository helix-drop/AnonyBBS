package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "/s2layout.html", nil)
}
