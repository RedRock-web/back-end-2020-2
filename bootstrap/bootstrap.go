package bootstrap

import (
	"back-end-2020-1/dao/redis"
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
	client, _ := redis.CreateClient()
	player := map[string]string{
		"a": "0",
		"b": "0",
		"c": "0",
		"d": "0",
		"e": "0",
	}
	client.HMSet("player", player)
}

//RouterInit 初始化 router
func RouterInit() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run()
}
