package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/handlers"
	"github.com/helix-drop/AnonyBBS/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./static")
	r.GET("/", handlers.Index)
	r.GET("/signin", handlers.Signin)
	r.GET("/signup", handlers.Signup)
	r.POST("handlers/register", handlers.Register)
	r.POST("/handlers/login", handlers.Login)
	r.GET("/handlers/showreply", handlers.ShowReply)
	//需要登录保护的
	auth := r.Group(``)
	auth.Use(middleware.AuthRequired())
	{
		auth.GET(`/logout`, handlers.Logout)
		auth.GET(`/mytopic`, handlers.MyTopic)
		auth.GET(`/myreply`, handlers.MyReply)
		auth.GET("/sendtopic", handlers.Stopic)
		auth.POST("/handlers/sendreply/", handlers.Sreply)
		auth.POST("/handlers/posttopic", handlers.PostTopic)
		auth.POST("/handlers/postreply", handlers.PostReply)
	}
	return r
}
