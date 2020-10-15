package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

// 查询社区id 与 社区名称
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows { // 查询结果为空
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (CommunityDetail *models.CommunityDetail, err error) {
	CommunityDetail = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time 
			   from community 
			   where community_id = ?`
	err = db.Get(CommunityDetail, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return CommunityDetail, err
}
