package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "wkb_comments/model"
	"wkb_comments/serializer"
)

type Page struct {
	Page    int `form:"page"`
	Size 	int `form:"size"`
	Desc	int `form:"desc"`
}

func Find(c *gin.Context) {
	id := c.Query("id")
	var comment Comment
	if err := DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": serializer.BuildComment(comment),
	})
}

func FindChildren(c *gin.Context) {
	id := c.Query("id")
	var comment Comment
	if err := DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"parent": serializer.BuildComment(comment),
		"children": serializer.BuildComments(comment.Children),
	})
}

func Index(c *gin.Context) {
	var comments []Comment
	var p Page
	if c.ShouldBindQuery(&p) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "参数错误",
		})
		return
	}
	page := p.Page
	pageSize := p.Size
	desc := p.Desc
	if page <= 0 {
		page = 1
	}
	var err error
	if desc == 1{	// 按时间反向
		err = DB.Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Find(&comments).Error
	} else {
		err = DB.Limit(pageSize).Offset((page-1)*pageSize).Find(&comments).Error
	}
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  err.Error(),
		})
	}
	var total int
	DB.Model(&Comment{}).Count(&total)
	var pageNum = total / pageSize
	if total % pageSize != 0{
		pageNum++
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"data": serializer.BuildComments(comments),
		"total": total,
		"page_num": pageNum,
	})
}
