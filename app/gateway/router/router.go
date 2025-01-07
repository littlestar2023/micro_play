package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	http "micro_play/app/gateway/handlers"
	"micro_play/app/gateway/middleware"
)

func InitRouters() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("my-session", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		v1.POST("/user/register", http.UserRegisterHandler)
		v1.POST("/user/login", http.UserLoginHandler)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("tasks", http.ListTaskHandler)
			authed.POST("task", http.CreateTaskHandler)
			authed.GET("task/:id", http.GetTaskHandler)       // task_id
			authed.PUT("task/:id", http.UpdateTaskHandler)    // task_id
			authed.DELETE("task/:id", http.DeleteTaskHandler) // task_id
		}
	}
	return ginRouter
}
