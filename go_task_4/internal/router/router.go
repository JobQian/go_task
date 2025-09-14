package router

import (
	"go_task_4/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(userhandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		// --- 用户认证 ---
		apiV1.POST("/register", userhandler.Register)
		apiV1.POST("/login", userhandler.Login)
	}

	return r
}
