package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/logger"
	"github.com/helix-drop/AnonyBBS/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterInfo struct {
	Name       string `bind:"required"`
	Email      string `bind:"required,email"`
	PassWord   string `bind:"required,min=6,max=20"`
	PwdConfirm string `bind:"eq=PassWord"`
}

func Register(c *gin.Context) {
	c.SetCookie("anony_id", "", -1, "/", "/", false, true)
	c.SetCookie("name", "", -1, "/", "/", false, true)
	var info RegisterInfo
	if err := c.ShouldBind(&info); err != nil {
		println(logger.Error)
		c.HTML(http.StatusOK, "/suc.html", gin.H{
			"msg": "Bind信息错误",
		})
		return
	}
	if IsExistEmail(info.Email) {
		println(logger.Error)
		c.HTML(http.StatusOK, "/suc.html", gin.H{
			"msg": "邮箱已被注册",
		})

		return
	}
	if IsExistName(info.Name) {
		println(logger.Error)
		c.HTML(http.StatusOK, "/suc.html", gin.H{
			"msg": "用户名已被注册",
		})
		return
	}
	if info.PassWord != info.PwdConfirm {
		println(logger.Error)
		c.HTML(http.StatusOK, "/suc.html", gin.H{
			"msg": "密码不匹配",
		})
		return
	}
	//加密
	bytesPwd, err := bcrypt.GenerateFromPassword([]byte(info.PassWord), 10)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "/suc.html", gin.H{
			"msg": "密码加密失败"})

		return
	}
	user := models.User{
		Name:     info.Name,
		Email:    info.Email,
		PassWord: string(bytesPwd),
	}
	//注册用户成功
	err = models.DB.Create(&user).Error
	if err != nil {
		logger.Error.Println("注册失败", err)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{
			"msg": "注册失败！",
		})
	} else {
		logger.Info.Println("注册成功", info.Name)
		c.HTML(http.StatusOK, "/suc.html", gin.H{
			"msg": "注册成功，欢迎!",
		})
	}
}
