package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TopicInfo struct {
	gorm.Model
	TopicTop     bool `gorm:"not null;default:false"`
	TopicTitle   string
	TopicContent string
	UserID       uint
}

type ReplyInfo struct {
	gorm.Model
	ReplyContent string
	ReplyFloorId uint `gorm:"AUTO_INCREMENT"`
	ReplyFloorTo uint
	UserID       uint
	TopicID      uint
}

func SendTopic(c *gin.Context) {
	return
}
func SendReply(c *gin.Context) {
	return
}
