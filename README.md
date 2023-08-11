# Tiktok

---

## 项目结构

```
|-- main.go           `主程序入口`
|-- go.mod
|-- go.sum
|
|-- controller        `Handle处理函数`
| |-- common.go
| |-- feed.go
|
|-- routers           `保存路由信息`
| |-- baseRouters.go
|
|-- util
| |-- author.go
| |-- init_db.go      `GORM配置文件`
| |-- tool.go         `工具类`
| |-- video.go
