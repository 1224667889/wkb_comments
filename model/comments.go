package model

import "github.com/jinzhu/gorm"

//Comment 评论模型
type Comment struct {
	gorm.Model
	Content       	string	`gorm:"not null"`	// 评论内容
	ParentId        uint	// 回复评论id
	UserName		string	`gorm:"not null"`	// 用户名
	ReplyName		string	// 回复名
	Avatar			string	`gorm:"default:'/static/img/avatar/index.jpg'"`// 头像
	TopicHash		string	`gorm:"size:255;not null;unique"`
	Children   		[]Comment `gorm:"ForeignKey:ParentId"`
}