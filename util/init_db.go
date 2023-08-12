package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const MYSQL_USERNAME = "root"
const MYSQL_PASSWORD = 1234
const MYSQL_HOST = "localhost"
const MYSQL_PORT = 3306

func init() {

	dsn := fmt.Sprintf("%s:%d@tcp(%s:%d)/tiktok?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT)

	// dsn := "root:1234@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
