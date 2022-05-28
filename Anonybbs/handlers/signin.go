package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signin(c *gin.Context) {

	//清除cookie
	c.SetCookie("anony_id", "", -1, "/", "/", false, true)
	c.SetCookie("name", "", -1, "/", "/", false, true)
	c.HTML(http.StatusOK, "/slayout.html", nil)
}
