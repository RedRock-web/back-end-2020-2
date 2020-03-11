package bootstrap

import (
	"back-end-2020-1/dao/dao_mysql"
	"back-end-2020-1/errors"
	"back-end-2020-1/router"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Init() 初始化项目
func Init() {
	MysqlInit()
	//RedisInit()
	RouterInit()
}

//MysqlInit() 初始化 dao_mysql 数据库
func MysqlInit() {
	var err error
	//gorm 不能创建数据库，需要自己建，可以用原生的 mysel，这里直接连 user 数据库了
	dao_mysql.G_db, err = dao_mysql.Connect(dao_mysql.DbLoginForm{"root", "root", "user"})
	errors.CheckError(err, "open dao_mysql dabatase user error!")
	//建 user 表

	if dao_mysql.G_db.HasTable(&dao_mysql.User{}) {
		fmt.Println("fwegewgewgewg")
		dao_mysql.G_db.AutoMigrate(&dao_mysql.User{})
	} else {
		err = dao_mysql.G_db.CreateTable(&dao_mysql.User{}).Error
		errors.CheckError(err, "create table named user error!")
	}
}

func RedisInit() {

}

//RouterInit 初始化 router
func RouterInit() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run()
}
