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
