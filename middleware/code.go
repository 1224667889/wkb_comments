package middleware

import (
	"github.com/gin-gonic/gin"
	. "wkb_comments/model"
)

//Code ApiCode验证中间件
func Code() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.DefaultPostForm("apiCode", "")
		var apiCode ApiCode
		DB.Where("Code = ?", code).First(&apiCode)
		if apiCode.UserName == "" {
			c.JSON(200, gin.H{
				"status": 404,
				"msg":  "apiCode不存在",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
