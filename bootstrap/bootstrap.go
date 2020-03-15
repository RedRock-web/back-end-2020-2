package bootstrap

import (
	"back-end-2020-1/dao"
	"back-end-2020-1/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Init() 初始化项目
func Init() {
	RedisInit()
	RouterInit()
}

//RedisInit 初始化 redis 数据库
//使用 hash 存储 player，key 为参赛选手名字，value 为票数
//使用 hash 存储 users，key 为用户 username, value 为 password
//使用 set 存储 status，存储参与投票的用户
func RedisInit() {
	dao.G_client, _ = dao.CreateClient()
	//使用 hash 初始化 5 位参赛选手
	player := map[string]string{
		"a": "0",
		"b": "0",
		"c": "0",
		"d": "0",
		"e": "0",
	}
	dao.G_client.HMSet("player", player)
}

//RouterInit 初始化 router
func RouterInit() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run()
}
