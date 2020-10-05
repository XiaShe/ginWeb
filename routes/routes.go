package routes

import (
	"net/http"
	"web_app/controllers"
	"web_app/logger"
	"web_app/settings"

	"github.com/gin-gonic/gin"
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

	// 用户登录
	r.POST("/login", controllers.LoginHandler)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	return r
}
