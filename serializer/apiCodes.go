package serializer

import (
	"wkb_comments/model"
)

type Code struct {
	ID       		uint   `json:"id"`
	UserName 		string `json:"UserName"`
	Code 			string `json:"Code"`
	CreateAt 		int64  `json:"create_at"`
}

//BuildCode 序列化评论
func BuildCode(code model.ApiCode) Code {
	return Code{
		ID:       code.ID,
		UserName: code.UserName,
		Code:	  code.Code,
		CreateAt: code.CreatedAt.Unix(),
	}
}