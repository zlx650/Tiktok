package models

type User struct {
	ID              int64
	Name            string `gorm:"column:name;type:varchar(32);unique"`
	FollowCount     uint   `gorm:"default:0;"`
	FollowerCount   uint   `gorm:"default:0;"`
	IsFollow        bool   `gorm:"default:false;"`
	Avatar          string `gorm:"type:varchar(255);"`
	BackgroundImage string `gorm:"type:varchar(255);"`
	Signature       string `gorm:"type:varchar(255);"`
	TotalFavorited  string `gorm:"type:varchar(255);"`
	WorkCount       uint   `gorm:"default:0;"`
	FavoriteCount   uint   `gorm:"default:0;"`
}

// TableName 表示配置操作数据库的表名称
func (User) TableName() string {
	return "user"
}

type UserForm struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token" binding:"required"`
}
