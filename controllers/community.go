package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（commutity_id， community_name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给客户端
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	communityID := c.Param("id") // 获取URL参数
	// 转换为int64
	id, err := strconv.ParseInt(communityID, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 查询到所有的社区（commutity_id， community_name）以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给客户端
		return
	}
	ResponseSuccess(c, data)
}
