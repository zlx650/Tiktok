package models

type User struct {
	ID              int64	 `json:"id"`
	Name            string `json:"name" gorm:"column:name;type:varchar(32);unique"`
	FollowCount     uint   `json:"follow_count" gorm:"default:0;"`
	FollowerCount   uint   `json:"follower_count" gorm:"default:0;"`
	IsFollow        bool   `json:"is_follow" gorm:"default:false;"`
	Avatar          string `json:"avatar" gorm:"type:varchar(255);"`
	BackgroundImage string `json:"background_image" gorm:"type:varchar(255);"`
	Signature       string `json:"signature" gorm:"type:varchar(255);"`
	TotalFavorited  string `json:"total_favorited" gorm:"type:varchar(255);"`
	WorkCount       uint   `json:"work_count" gorm:"default:0;"`
	FavoriteCount   uint   `json:"favorite_count" gorm:"default:0;"`
}

// TableName 表示配置操作数据库的表名称
func (User) TableName() string {
	return "user"
}

type UserForm struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token" binding:"required"`
}
