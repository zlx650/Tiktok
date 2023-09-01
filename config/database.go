package config

import "os"

// 1024code 中从环境变量获取值
var MYSQL_USERNAME = os.Getenv("MYSQL_USER")
var MYSQL_PASSWORD = "JAAquZNn"
var MYSQL_HOST = os.Getenv("MYSQL_HOST")
var MYSQL_PORT = os.Getenv("MYSQL_PORT")

// var MYSQL_USERNAME = "root"
// var MYSQL_PASSWORD = "kRqqJTjJ"
// var MYSQL_HOST = "172.16.32.66"
// var MYSQL_PORT = "50333"
