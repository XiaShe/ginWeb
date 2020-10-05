package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

// 存放业务逻辑代码

// SignUp 用户注册验证
func SignUp(p *models.ParamSignUp) (err error) {

	// 1. 判断用户是否存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	// 2. 生成UID
	userID := snowflake.GenID()
	// 构造一个Users实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 3. 保存到数据库
	return mysql.InsertUser(user)
}

// SignIn 用户登录验证
func Login(p *models.ParamLogin) (err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)

}
