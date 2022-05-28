package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/models"
	"net/http"
)

func MyReply(c *gin.Context) {
	reply := make([]models.Reply, 10)
	name, err := c.Cookie("name")
	if err != nil {
		println("Cookie无效！")
	}
	models.DB.Where("poster=?", name).Order("created_at Desc").Find(&reply)
	c.HTML(http.StatusOK, "/m2layout.html", reply)
}
