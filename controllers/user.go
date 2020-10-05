package controllers

import (
	"fmt"
	"log"
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		// 翻译错误
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 将放回的错误格式化
		})
		return
	}

	// 2. 业务处理
	err = logic.SignUp(p)
	if err != nil {
		fmt.Println(err.Error())
		log.Println("user业务处理没啥错误")
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "Sign up success",
	})

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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		// 翻译错误
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 将放回的错误格式化
		})
		return
	}

	// 2. 业务处理
	err = logic.Login(p)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "Login success",
	})

}
