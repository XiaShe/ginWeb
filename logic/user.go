package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
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
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 登录成功
	err = mysql.Login(user) // user传递的为一个指针
	if err != nil {
		return "", err
	}
	// 根据登录的 用户id 和 用户名 与其它默认参数 生成 JWT token
	return jwt.GenToken(user.UserID, user.Username)

}
