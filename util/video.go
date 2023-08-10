package util

type Video struct {
	Id            int64  `json:"id"`
	AuthorId      int64  `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    int8   `json:"is_favorite"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
}

// TableName 表示配置操作数据库的表名称
func (Video) TableName() string {
	return "video"
}
