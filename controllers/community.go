package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
)

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（commutity_id， community_name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList()", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给客户端
		return
	}
	ResponseSuccess(c, data)
}
