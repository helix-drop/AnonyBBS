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
	var info RegisterInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		println(logger.Error)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "json信息错误"})
		return
	}
	if !(IsExistName(info.Name) && IsExistEmail(info.Email)) {
		println(logger.Error)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户名或邮箱已注册"})
		return
	}
	if info.PassWord != info.PwdConfirm {
		println(logger.Error)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码不匹配"}) //TODO 具体化错误信息
		return
	}
	bytesPwd, err := bcrypt.GenerateFromPassword([]byte(info.PassWord), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "密码加密失败"})

		return
	}
	user := models.User{
		Name:     info.Name,
		Email:    info.Email,
		PassWord: string(bytesPwd),
	}

	err = models.DB.Create(&user).Error
	if err != nil {
		logger.Error.Println("注册失败", err)
	}
	logger.Info.Println("注册成功", info.Name)
}
