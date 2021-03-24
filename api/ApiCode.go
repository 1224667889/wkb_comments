package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wkb_comments/model"
	"wkb_comments/serializer"
)

type NewCode struct {
	ApiCode     	string 	`form:"apiCode" json:"apiCode" xml:"apiCode"  binding:"required"`
	UserName    	string 	`form:"username" json:"username" xml:"username"  binding:"required"`
}

func CreateCode(c *gin.Context) {
	var code NewCode
	if err := c.ShouldBind(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCode := model.ApiCode{
		Code:   	code.ApiCode,
		UserName:  	code.UserName,
	}
	var apiCode model.ApiCode
	model.DB.Where("Code = ?", newCode.Code).First(&apiCode)
	if  apiCode.Code == ""{
		model.DB.Create(&newCode)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": serializer.BuildCode(newCode),
		})
	}else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "have this code",
		})
	}
}

func RemoveCode(c *gin.Context) {
	code := c.Query("code")
	var apiCode model.ApiCode
	if err := model.DB.Where("Code = ?", code).First(&apiCode).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "not found",
		})
		return
	}
	//r := map[string]string{
	//	"username": comment.UserName,
	//	"content": comment.Content,
	//}
	if err := model.DB.Delete(&apiCode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "something wrong when delete",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": serializer.BuildCode(apiCode),
	})
}