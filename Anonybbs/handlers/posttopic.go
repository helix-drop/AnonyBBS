package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/logger"
	"github.com/helix-drop/AnonyBBS/models"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type TopicInfo struct {
	TopicTitle   string
	TopicContent string
	AnonyId      string
}

func PostTopic(c *gin.Context) {
	var post TopicInfo
	if err := c.ShouldBind(&post); err != nil {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "json信息错误", "backto": "/"})
		return
	}
	if post.TopicTitle == "" {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "标题不能为空！", "backto": "/"})
		return
	}
	if post.TopicContent == "" {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "内容不能为空！", "backto": "/"})
		return
	}
	name, err := c.Cookie("name")
	if err != nil {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "登录状态错误！", "backto": "/"})
		return
	}
	anonyid, err := c.Cookie("anony_id")
	if err != nil {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "登录状态错误！", "backto": "/"})
		return
	}
	//获取user表中，name对应的id
	var user models.User
	models.DB.Where("name=?", name).First(&user)
	topic := models.Topic{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		TopicTitle:   post.TopicTitle,
		TopicContent: post.TopicContent,
		Poster:       name,
		UserID:       user.ID,
		AnonyId:      anonyid,
	}

	models.DB.Create(&topic)
	reply := models.Reply{
		TContent: post.TopicTitle,
		TTitle:   post.TopicContent,
		TopicID:  topic.ID,
		UserID:   user.ID,
	}
	err = models.DB.Create(&reply).Error
	if err != nil {
		logger.Error.Println("发帖失败！", err)
		c.HTML(http.StatusMethodNotAllowed, "/suc.html", gin.H{
			"msg":    "发帖失败！",
			"backto": "/",
		})
	} else {
		logger.Info.Println("发帖成功！")
		c.HTML(http.StatusOK, "/suc.html", gin.H{
			"msg":    "发帖成功！",
			"backto": "/",
		})
	} //创建一条id为0的回复

}
