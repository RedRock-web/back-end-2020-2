package bootstrap

import (
	"back-end-2020-1/dao/dao_my"
	"back-end-2020-1/errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Init() 初始化项目
func Init() {
	MysqlInit()
	//RedisInit()
	//RouterInit()
}

//MysqlInit() 初始化 dao_my 数据库
func MysqlInit() {
	var err error
	//gorm 不能创建数据库，需要自己建，可以用原生的 mysel，这里直接连 user 数据库了
	dao_my.G_db, err = dao_my.Connect(dao_my.DbLoginForm{"root", "root", "user"})
	errors.CheckError(err, "open dao_my dabatase user error!")
	//建 user 表

	if dao_my.G_db.HasTable(&dao_my.User{}) {
		fmt.Println("fwegewgewgewg")
		dao_my.G_db.AutoMigrate(&dao_my.User{})
	} else {
		err = dao_my.G_db.CreateTable(&dao_my.User{}).Error
		errors.CheckError(err, "create table named user error!")
	}
}

func RedisInit() {

}

func RouterInit() {

}
