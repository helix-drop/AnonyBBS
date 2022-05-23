package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/logger"
	"github.com/helix-drop/AnonyBBS/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginInfo struct {
	Email    string `bind:"required,email"`
	PassWord string `bind:"required,min=6,max=20"`
}

func Login(c *gin.Context) {
	var info LoginInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		println(logger.Error)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "json信息错误"})
		return
	}
	user := models.User{}
	if IsExistEmail(info.Email) {
		println(logger.Error)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户不存在，请注册！"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(info.PassWord)); err != nil {
		println(logger.Error)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码错误！"})
		return
	}
	AnonyID := GenerateRandomString()
	c.SetCookie("anony_id", AnonyID, 7200, "/", "/", false, true)
	logger.Info.Println("登录成功", user.Name)
	c.JSON(http.StatusOK, gin.H{"msg": "登录成功!", "data": user.Name})

}
