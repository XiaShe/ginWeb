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
