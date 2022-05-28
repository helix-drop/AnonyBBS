package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/helix-drop/AnonyBBS/models"
	"net/http"
	"strconv"
)

type uri struct {
	ID string
}

/*
这个函数无意义增加工作量，舍弃
type output struct {
	Topicid       uint
	Topictitle    string
	Toppiccontent string
	Replyfloorid  []uint
	Replyfloorto  []uint
	Replycontent  []string
	Poster        []string
	Createdat     []time.Time
}

func ShowReply(c *gin.Context) {
	var info uri
	if err := c.ShouldBind(&info); err != nil {
		println("请求错误！")
	}
	i, err := strconv.ParseUint(info.ID, 10, 0)
	if err != nil {
		println("请求参数错误！")
	}
	replys := make([]models.Reply, 10)
	models.DB.Where("topic_id=?", i).Order("created_at Desc").Find(&replys)
	topic := models.Topic{}
	models.DB.Where("ID=?", i).Order("created_at Desc").Find(&topic)

	rfi := make([]uint, 10)
	for k, _ := range replys {
		rfi[k] = replys[k].ReplyFloorId
	}
	rft := make([]uint, 10)
	for k, _ := range replys {
		rft[k] = replys[k].ReplyFloorTo
	}
	rfc := make([]string, 10)
	for k, _ := range replys {
		rfc[k] = replys[k].ReplyContent
	}
	aid := make([]string, 10)
	for k, _ := range replys {
		aid[k] = replys[k].AnonyId
	}

	ca := make([]time.Time, 10)
	for k, _ := range replys {
		ca[k] = replys[k].CreatedAt
	}
	put := output{
		Topicid:       uint(i),
		Topictitle:    topic.TopicTitle,
		Toppiccontent: topic.TopicContent,
		Replyfloorid:  rfi,
		Replyfloorto:  rft,
		Replycontent:  rfc,
		Poster:        aid,
		Createdat:     ca,
	}

	c.HTML(http.StatusOK, "/rlayout.html", put)
}

*/
func FindTopicInfo(r models.Reply) (topictitle, topiccontent string) {
	topic := models.Topic{}
	models.DB.Where("ID=?", r.TopicID).First(&topic)
	topictitle = topic.TopicTitle
	topiccontent = topic.TopicContent
	return topictitle, topiccontent
}

func ShowReply(c *gin.Context) {
	var info uri
	if err := c.ShouldBind(&info); err != nil {
		println("请求错误！")
	}
	i, err := strconv.ParseUint(info.ID, 10, 0)
	if err != nil {
		println("请求参数错误！")
	}
	replys := make([]models.Reply, 10)
	models.DB.Where("topic_id=?", i).Order("created_at Asc").Find(&replys)
	topic := models.Topic{}
	models.DB.Where("ID=?", i).Order("created_at Desc").Find(&topic)
	replys[0].TTitle = topic.TopicTitle
	replys[0].TContent = topic.TopicContent
	c.HTML(http.StatusOK, "/rlayout.html", replys)

}
