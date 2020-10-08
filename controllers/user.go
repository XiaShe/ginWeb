package controllers

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// 处理注册请求
func SignUpHandler(c *gin.Context) {

	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	err := c.ShouldBindJSON(p) // 参数绑定

	if err != nil {
		zap.L().Error("SingUpHand with invalid param", zap.Error(err)) // 向日志中记录错误信息

		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		// 将具体错误信息翻译为中文
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 业务处理
	err = logic.SignUp(p)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserExist) { // 如果错误为ErrorUserNotExist
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)

}

// 处理登录请求
func LoginHandler(c *gin.Context) {

	// 1. 获取登录时的参数
	p := new(models.ParamLogin)
	err := c.ShouldBindJSON(p) // 参数绑定

	if err != nil {
		zap.L().Error("LoginHand with invalid param", zap.Error(err)) // 向日志中记录错误信息

		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		// 将具体错误信息翻译为中文
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 业务处理, 获得 token
	token, err := logic.Login(p)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) { // 如果错误为ErrorUserNotExist
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, token)

}
