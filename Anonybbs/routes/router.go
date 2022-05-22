package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/handlers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	return r
}
