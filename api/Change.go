package api

import (
	"github.com/gin-gonic/gin"
)

func Change(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "there is nothing",
	})
}