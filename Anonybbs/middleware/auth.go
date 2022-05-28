package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//验证登录状态
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Request.Cookie("name")
		if cookie == nil {
			c.HTML(http.StatusUnauthorized, "/suc.html", gin.H{"msg": "请先登陆"})
			c.Abort()
		}
		c.Next()
	}
}
