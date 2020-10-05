package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

// 把每一步数据库操作封装成函数
// 等待logic层根据业务需求调用

const secret = "xiashe"

// CheckUserExist  检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return errors.New("用户不存在")
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func Login(user *models.User) (err error) {
	oPassword := user.Password // 登录时输入的密码
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	// 判断用户名是否存在（是否no rows）
	if err == sql.ErrNoRows {
		return errors.New("用户名不存在")
	}

	if err != nil {
		return err // 查询数据库失败
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)

	if password != user.Password {
		return errors.New("密码错误")
	}

	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密（数据库不能储存明文密码）
	user.Password = encryptPassword(user.Password)
	// 入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// 加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
