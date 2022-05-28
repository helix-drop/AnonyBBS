package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//发帖
func Stopic(c *gin.Context) {
	c.HTML(http.StatusOK, "/tlayout.html", nil)
}

//回复
type rinfo struct {
	TopicID    uint
	ReplyFloor uint
}

func Sreply(c *gin.Context) {
	var r rinfo
	_ = c.ShouldBind(&r)

	c.HTML(http.StatusOK, "/t2layout.html", r)

}
