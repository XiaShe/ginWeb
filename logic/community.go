package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据， 查找所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	// 根据ID查询不同社区（community）详情
	return mysql.GetCommunityDetailByID(id)
}
