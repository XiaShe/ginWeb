package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据， 查找所有的community 并返回
	return mysql.GetCommunityList()
}
