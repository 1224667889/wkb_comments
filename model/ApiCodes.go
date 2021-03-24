package model

import "github.com/jinzhu/gorm"

//ApiCode api身份识别模型
type ApiCode struct {
	gorm.Model
	Code	       	string	`gorm:"not null;unique"`	// ApiCode(唯一)
	UserName		string	`gorm:"not null"`			// 申请用户名
}
