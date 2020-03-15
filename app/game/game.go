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
	fmt.Println(account.G_username)
	flag, _ := dao.G_client.SIsMember("status", account.G_username).Result()
	return flag
}

func Vote(c *gin.Context) {
	targe := c.PostForm("targe")
	fmt.Println(targe)
	if TargeIsLegal(targe) {
		if HaveVote(targe) {
			response.Error(c, 10010, "had vote!")
		} else {
			dao.G_client.SAdd(targe+"Votes", account.G_username)
		}
	} else {
		response.FormError(c)
	}
}

func CancelVote(c *gin.Context) {
	targe := c.PostForm("targe")
	if TargeIsLegal(targe) {
		if HaveVote(targe) {
			dao.G_client.SRem(targe+"Votes", account.G_username)
		} else {
			response.Error(c, 10010, "have not vote!")
		}
	} else {
		response.FormError(c)
	}
}

func TargeIsLegal(targe string) bool {
	return targe == "a" || targe == "b" || targe == "c" || targe == "d" || targe == "e"
}

func HaveVote(targe string) (flag bool) {
	flag, _ = dao.G_client.SIsMember(targe+"Votes", account.G_username).Result()
	return flag
}
