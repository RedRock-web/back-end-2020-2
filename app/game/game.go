package game

import (
	"back-end-2020-1/app/account"
	"back-end-2020-1/dao"
	"back-end-2020-1/errors"
	"back-end-2020-1/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Enter(c *gin.Context) {
	if HaveEnter() {
		fmt.Println("###")
		response.Error(c, 10007, "aready have enter game!")
	} else {
		fmt.Println("SDfse")
		err := dao.G_client.SAdd("status", account.G_username).Err()
		errors.CheckError(err, "enter game error!")
	}
}

func Retire(c *gin.Context) {
	if HaveEnter() {
		err := dao.G_client.SRem("status", account.G_username).Err()
		errors.CheckError(err, "enter game error!")
	} else {
		response.Error(c, 10008, "Did not participate in the competition")
	}
}

func HaveEnter() bool {
	flag, _ := dao.G_client.SIsMember("status", account.G_username).Result()
	return flag
}
