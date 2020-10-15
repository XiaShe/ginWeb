package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成 post ID
	p.ID = int64(snowflake.GenID())

	// 2. 保存到数据库
	return mysql.CreatePost(p)
}
