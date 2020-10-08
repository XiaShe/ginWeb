package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	// 用户注册
	r.POST("/signup", controllers.SignUpHandler)

	// 用户登录（JWT协议）
	r.POST("/login", controllers.LoginHandler)

	// 用户请求，JWTAuthMiddleware 中间件 会判断 请求头 中是否有有效的 JWT token
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 请求头存在有效token方能进行下一步操作
		c.String(http.StatusOK, "目前处于登录状态")

	})
	return r
}
