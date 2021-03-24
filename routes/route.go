package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"wkb_comments/api"
	"wkb_comments/middleware"
)

//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("******"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		CommentGroup := v1.Group("/comment")
		{
			CommentGroup.GET("", api.Find)
			CommentGroup.GET("/children", api.FindChildren)
			CommentGroup.GET("/index", api.Index)
			CommentGroup.POST("", middleware.Code(), api.Add)
			CommentGroup.DELETE("", middleware.JWT(), api.Del)
			CommentGroup.PUT("", api.Change)
		}

		adminGroup := v1.Group("/admin")
		adminGroup.Use(middleware.JWT())
		{
			adminGroup.POST("/apiCode", api.CreateCode)
			adminGroup.DELETE("/apiCode", api.RemoveCode)
		}

	}
	return r
}


