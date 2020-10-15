package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的校验
	p := new(models.Post)
	err := c.ShouldBindJSON(p)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 从 c 获取当前发送请求的用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	// 2. 创建帖子
	err = logic.CreatePost(p)
	if err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}
