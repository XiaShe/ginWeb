package models

import "time"

// 内存对齐：相同类型尽量放在一起 <此处就不弄了>
type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	AuthorID    int64     `json:"author_id" db:"auther_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"` // 由于mysql版本问题，默认创建时间为 2016-11-01 08:10:10
}
