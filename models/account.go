package models

type Account struct {
	ID       int64
	UserId   int64
	Username string `gorm:"varchar(32);unique;not null" binding:"required"`
	Password string `gorm:"varchar(32);not null" binding:"required"`
}

// TableName 表示配置操作数据库的表名称
func (Account) TableName() string {
	return "account"
}

