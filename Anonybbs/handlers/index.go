package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/models"
	"net/http"
)

func Index(c *gin.Context) {
	topics := make([]models.Topic, 10)
	models.DB.Order("created_at Desc").Find(&topics)
	c.HTML(http.StatusOK, "/layout.html", topics)
}
