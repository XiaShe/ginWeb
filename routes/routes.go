package routes

import (
	"github.com/gin-gonic/gin"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middlewares"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// 添加中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("api/v1")

	// 用户注册
	v1.POST("/signup", controllers.SignUpHandler)

	// 用户登录（JWT协议）
	v1.POST("/login", controllers.LoginHandler)

	// 应用中间件
	/*
		JWTAuthMiddleware中间件 会判断请求头中是否有有效的 JWT token
		请求头存在有效token方能进行下一步操作
	*/
	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controllers.CommunityHandler)           // 社区列表详情
		v1.GET("/community/:id", controllers.CommunityDetailHandler) // 不同社区详情

		v1.POST("/post", controllers.CreatePostHandler)       // 创建帖子
		v1.GET("/post/:id", controllers.GetPostDetailHandler) // 获得帖子详情信息
		v1.GET("/posts/", controllers.GetPostListHandler)
	}

	return r
}
