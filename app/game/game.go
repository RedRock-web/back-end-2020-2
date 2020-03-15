package game

import (
	"back-end-2020-1/app/account"
	"back-end-2020-1/dao/dao_mysql"
	"back-end-2020-1/errors"
	"back-end-2020-1/response"
	"github.com/gin-gonic/gin"
)

func Enter(c *gin.Context) {
	if HaveEnter() {
		response.Error(c, 10007, "aready have enter game!")
	} else {
		user := dao_mysql.User{Username: account.G_username}
		err := dao_mysql.G_db.Model(&user).Update("participation_status", 1).Error
		errors.CheckError(err, "enter game error!")
	}
}

func Retire(c *gin.Context) {
	if HaveEnter() {
		user := dao_mysql.User{Username: account.G_username}
		err := dao_mysql.G_db.Model(&user).Update("participation_status", 0).Error
		errors.CheckError(err, "retire game error!")
	} else {
		response.Error(c, 10008, "Did not participate in the competition")
	}
}

func HaveEnter() bool {
	var u dao_mysql.User

	dao_mysql.G_db.Where(dao_mysql.User{
		Username: account.G_username,
	}).First(&u)

	return u.Participation_status == 1
}
