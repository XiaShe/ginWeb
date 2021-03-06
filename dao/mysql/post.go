package mysql

import "web_app/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post( 
			   post_id, title, content, author_id, community_id)
			   value(?, ?, ?, ?, ?)`

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
			   from post
			   where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

// 显示的页数是page，每页显示的条目数为size
func GetPostList(page, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
			   from post 
			   limit ?, ?`
	post = make([]*models.Post, 0, 2)
	err = db.Select(&post, sqlStr, (page-1)*size, size)
	if err != nil {
		return nil, err
	}
	return
}
