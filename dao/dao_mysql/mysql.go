package dao_mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//User struct 用于创建 users 表，表示一个 user
//Participation_status 表示参赛状态，默认为 0, 0 表示未参赛，1 表示参赛
//Vote_num 表示获得投票数，默认为 0
type User struct {
	gorm.Model
	Username             string
	Password             string
	Participation_status int `gorm:"default:0"`
	Vote_num             int `gorm:"default:0"`
}

//G_db 用于存储 dao_mysql 数据库
var G_db *gorm.DB

type DbLoginForm struct {
	Username string
	Password string
	DbName   string
}

func Connect(form DbLoginForm) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", form.Username+":"+form.Password+"@(127.0.0.1:3306)/"+form.DbName+"?charset=utf8&parseTime=true")
	return db, err
}

func Insert(i interface{}, errMsg string) error {
	u := i.(User)
	if err := G_db.Create(&u).Error; err != nil {
		errors.New(errMsg)
		return err
	}
	return nil
}
