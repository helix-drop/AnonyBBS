package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"index"`
	Email    string `gorm:"index"`
	PassWord string
	Topics   []Topic
	Replies  []Reply
}

type Topic struct {
	gorm.Model
	TopicTop     bool `gorm:"not null;default:false"`
	TopicTitle   string
	TopicContent string
	UserID       uint
	Replies      []Reply
}

type Reply struct {
	gorm.Model
	ReplyContent string
	ReplyFloorId uint `gorm:"AUTO_INCREMENT"`
	ReplyFloorTo uint
	UserID       uint
	TopicID      uint
}
