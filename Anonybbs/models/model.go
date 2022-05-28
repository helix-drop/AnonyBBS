package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"index"`
	Email    string `gorm:"index;email;unique"`
	PassWord string
	Topics   []Topic
	Replies  []Reply
}

type Topic struct {
	gorm.Model
	TopicTop     bool
	TopicTitle   string
	TopicContent string
	UserID       uint
	Poster       string
	AnonyId      string
	Replies      []Reply
}

type Reply struct {
	gorm.Model
	ReplyContent string
	ReplyFloorId uint
	ReplyFloorTo uint
	Poster       string
	UserID       uint
	AnonyId      string
	TID          uint
	TopicID      uint
	TTitle       string
	TContent     string
}
