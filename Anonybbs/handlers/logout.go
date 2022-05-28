package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Logout(c *gin.Context) {
	// 设置cookie过期
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookie := http.Cookie{Name: "userID", Value: "", Expires: expiration}
	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, gin.H{"msg": "退出成功"})
}
