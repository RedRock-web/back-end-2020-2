package game

import (
	"back-end-2020-1/app/account"
	"back-end-2020-1/dao"
	"back-end-2020-1/errors"
	"back-end-2020-1/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"strconv"
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
	if TargeIsLegal(targe) {
		if HaveVote(targe) {
			response.Error(c, 10010, "had vote!")
		} else {
			dao.G_client.SAdd(targe+"Votes", account.G_username)
			dao.G_client.HIncrBy("player", targe, 1)
			dao.G_client.HIncrBy("user_vote_num", targe, 1)
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
			dao.G_client.HIncrBy("player", targe, -1)
			dao.G_client.HIncrBy("user_vote_num", targe, 1)
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

type Pair struct {
	name string
	num  int
}

type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].num < p[j].num }
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func GetBoard(c *gin.Context) {
	playerStatus := make(map[string]int, 10)

	aVoteNum, _ := dao.G_client.HGet("player", "a").Result()
	fmt.Println(aVoteNum)
	playerStatus["a"], _ = strconv.Atoi(aVoteNum)

	bVoteNum, _ := dao.G_client.HGet("player", "b").Result()
	playerStatus["b"], _ = strconv.Atoi(bVoteNum)

	cVoteNum, _ := dao.G_client.HGet("player", "c").Result()
	playerStatus["c"], _ = strconv.Atoi(cVoteNum)

	dVoteNum, _ := dao.G_client.HGet("player", "d").Result()
	playerStatus["d"], _ = strconv.Atoi(dVoteNum)

	eVoteNum, _ := dao.G_client.HGet("player", "e").Result()
	playerStatus["e"], _ = strconv.Atoi(eVoteNum)
	p := sortMapByValue(playerStatus)

	data := []gin.H{}
	for k, v := range p {
		data = append(data, gin.H{
			"rank":     k,
			"name":     v.name,
			"vote_num": v.num,
		})
	}
	response.OkWithData(c, data)
}

func CanVote() bool {
	tempNum, _ := dao.G_client.HGet("user_vote_num", account.G_username).Result()
	num, _ := strconv.Atoi(tempNum)
	return num <= 3 && num >= 0
}

func RestoreVoteNum() {
	dao.G_client.HMSet("user_vote_num", map[string]string{"a": "0", "b": "0", "c": "0", "d": "0", "e": "0"})
}
