package serializer

import (
	"wkb_comments/model"
)

type Comment struct {
	ID       		uint   `json:"id"`
	UserName 		string `json:"UserName"`
	ReplyName 		string `json:"ReplyName"`
	Avatar   		string `json:"Avatar"`
	ParentId   		uint   `json:"ParentId"`
	CreateAt 		int64  `json:"create_at"`
}

//BuildComment 序列化评论
func BuildComment(comment model.Comment) Comment {
	return Comment{
		ID:       comment.ID,
		UserName: comment.UserName,
		ReplyName:comment.ReplyName,
		Avatar:	  comment.Avatar,
		ParentId: comment.ParentId,
		CreateAt: comment.CreatedAt.Unix(),
	}
}

func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}