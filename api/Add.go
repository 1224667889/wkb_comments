package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wkb_comments/model"
	"wkb_comments/serializer"
)

type NewComment struct {
	Content     string 	`form:"content" json:"content" xml:"content"  binding:"required"`
	ApiCode     string 	`form:"apiCode" json:"apiCode" xml:"apiCode"  binding:"required"`
	UserName    string 	`form:"username" json:"username" xml:"username"  binding:"required"`
	TopicHash   string  `form:"topicHash" json:"topicHash" xml:"topicHash"  binding:"required"`
	ReplyName   string 	`form:"reply" json:"reply" xml:"reply"`
	ParentId 	uint 	`form:"parentId" json:"parentId" xml:"parentId"`
	Avatar		string  `form:"avatar" json:"avatar" xml:"avatar binding:"default:'/static/img/avatar/index.jpg'"`
}


func Add(c *gin.Context) {
	var comment NewComment
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment := model.Comment{
		Content:   comment.Content,
		ParentId:  comment.ParentId,
		UserName:  comment.UserName,
		ReplyName: comment.ReplyName,
		Avatar:    comment.Avatar,
		TopicHash: comment.TopicHash,
	}

	if model.DB.NewRecord(newComment) {
		model.DB.Create(&newComment)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": serializer.BuildComment(newComment),
		})
	}else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "have this comment",
		})
	}
}