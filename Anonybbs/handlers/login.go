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

	//接收登录信息
	var info LoginInfo
	if err := c.ShouldBind(&info); err != nil {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/slayout.html", gin.H{"msg": "bind信息错误"})
		return
	}
	user := models.User{}
	models.DB.First(&user)
	if !IsExistEmail(info.Email) {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/slayout.html", gin.H{"msg": "用户不存在，请注册！"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(info.PassWord)); err != nil {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/slayout.html", gin.H{"msg": "密码错误！"})
		return
	}
	name := models.User{}
	models.DB.Where("Email=?", info.Email).First(&name)
	//生成匿名ID
	AnonyID := GenerateRandomString()
	c.SetCookie("anony_id", AnonyID, 7200, "/", "/", false, true)
	c.SetCookie("name", name.Name, 7200, "/", "/", false, true)
	logger.Info.Println("登录成功", name.Name)
	c.HTML(http.StatusOK, "/suc.html", gin.H{"msg": "登录成功， " + name.Name + "用户您好! " + "您的匿名ID为: " + AnonyID})

}
