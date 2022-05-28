package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/models"
	"net/http"
)

func MyTopic(c *gin.Context) {
	topics := make([]models.Topic, 10)
	name, err := c.Cookie("name")
	if err != nil {
		println("Cookie无效！")
	}
	models.DB.Where("poster=?", name).Order("created_at Desc").Find(&topics)
	c.HTML(http.StatusOK, "/mlayout.html", topics)
}
