package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "wkb_comments/model"
	"wkb_comments/serializer"
)

func Del(c *gin.Context) {
	id := c.Query("id")
	var comment Comment
	if err := DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "not found",
		})
		return
	}
	//r := map[string]string{
	//	"username": comment.UserName,
	//	"content": comment.Content,
	//}
	if err := DB.Delete(comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "something wrong when delete",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": serializer.BuildComment(comment),
	})
}