package router

import (
	"go_task_4/internal/handler"
	"go_task_4/pkg/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(userhandler *handler.UserHandler, posthandler *handler.PostHandler, commentnandler *handler.CommentHandler) *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		// --- 用户认证 ---
		apiV1.POST("/register", userhandler.Register)
		apiV1.POST("/login", userhandler.Login)

		apiV1.GET("/allposts", posthandler.FindAllPost)
		apiV1.GET("/postbyid/:id", posthandler.GetByID)
		apiV1.GET("postsbyuserid/:id", posthandler.GetByUserID)

		apiV1.GET("/commentbyid/:id", commentnandler.GetByID)
		apiV1.GET("commentsbyuserid/:id", commentnandler.GetByUserID)
		apiV1.GET("commentsbypostid/:id", commentnandler.GetByPostID)

	}

	auth := apiV1.Group("/")
	auth.Use(utils.JWTAuthMiddleware())
	{
		auth.POST("/post", posthandler.CreatePost)
		auth.PUT("/post/:id", posthandler.UpdatePost)
		auth.DELETE("/post/:id", posthandler.DeletePost)

		auth.POST("/comment", commentnandler.CreateComment)
		auth.PUT("/comment/:id", commentnandler.UpdateComment)
		auth.DELETE("/comment/:id", commentnandler.DeleteComment)
	}
	return r
}
