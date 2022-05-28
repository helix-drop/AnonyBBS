package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/logger"
	"github.com/helix-drop/AnonyBBS/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type ReplyInfo struct {
	ReplyContent string
	ReplyFloorId uint `gorm:"AUTO_INCREMENT"`
	ReplyFloor   string
	Poster       string
	TopicID      uint
	AnonyId      string
}

func PostReply(c *gin.Context) {
	var post ReplyInfo
	if err := c.ShouldBind(&post); err != nil {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "json信息错误"})
		return
	}
	if post.ReplyContent == "" {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "内容不能为空！"})
		return
	}
	RID, err := strconv.ParseUint(post.ReplyFloor, 10, 0)
	if err != nil {
		println("请求参数错误！")
	}
	name, err := c.Cookie("name")
	if err != nil {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "登录状态错误！"})
		return
	}
	anonyid, err := c.Cookie("anony_id")
	if err != nil {
		println(logger.Error)
		c.HTML(http.StatusBadRequest, "/suc.html", gin.H{"msg": "登录状态错误！"})
		return
	}
	//获取user表中，name对应的id
	var user models.User
	var topic models.Topic
	models.DB.Where("name=?", name).First(&user)
	models.DB.Where("id=?", post.TopicID).First(&topic)
	post.ReplyFloorId = uint(len(topic.Replies) + 1)
	reply := models.Reply{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ReplyContent: post.ReplyContent,
		ReplyFloorId: post.ReplyFloorId,
		ReplyFloorTo: uint(RID),
		Poster:       name,
		UserID:       user.ID,
		AnonyId:      anonyid,
		TopicID:      post.TopicID,
	}
	a := int(post.TopicID)
	id := strconv.Itoa(a)
	err = models.DB.Create(&reply).Error
	if err != nil {
		logger.Error.Println("回帖失败！", err)
	} else {
		logger.Info.Println("回帖成功！")
		c.HTML(http.StatusOK, "/suc.html", gin.H{"msg": "回帖成功！", "backto": "handlers/showreply?ID=" + id})
	}
}
