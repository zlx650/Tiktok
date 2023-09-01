package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiktok/config"
)

var DB *gorm.DB
var err error

func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/tiktok?charset=utf8mb4&parseTime=True&loc=Local", config.MYSQL_USERNAME, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT)

	// dsn := "root:1234@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	// fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	// 更新数据库中videoURL数据
	err = UpdateVideoURL()

	if err != nil {
		return err
	}

	return nil
}
