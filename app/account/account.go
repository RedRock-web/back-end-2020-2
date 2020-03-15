package account

import (
	"back-end-2020-1/app/jwts"
	"back-end-2020-1/config"
	"back-end-2020-1/dao"
	"back-end-2020-1/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

var G_username string

func Register(c *gin.Context) {
	f := config.LoginForm{}

	if err := c.ShouldBindJSON(&f); err != nil {
		errors.New("bind json error!")
		response.FormError(c)
	} else if IsRegiste(f.Username) {
		response.Error(c, 10003, "user exist!")
	} else {
		dao.G_client.HSet("users", f.Username, f.Password)
		fmt.Println("ppppppppppppp")
		fmt.Println(f)
		token := GetJwt(f, "register create jwt error!")
		c.SetCookie("auth", token, 1000, "/", "127.0.0.1:8080", false, true)
		G_username = f.Username
	}
}

func IsRegiste(username string) bool {
	_, err := dao.G_client.HGet("users", username).Result()
	return err == nil
}

func GetJwt(f config.LoginForm, errMsg string) string {
	j := jwts.NewJwt()
	token, err := j.Create(f, "redrock")
	if err != nil {
		errors.New(errMsg)
	}
	return token
}

func Login(c *gin.Context) {
	f := config.LoginForm{}
	token := GetJwt(f, "error")

	if IsLogin(c) {
		response.OkWithData(c, "aready login!")
	} else if err := c.ShouldBindJSON(&f); err != nil {
		errors.New("bind json error!")
		response.FormError(c)
	} else {
		if !IsRegiste(f.Username) {
			response.Error(c, 10005, "unregistered!")
		} else if !PasswdIsOk(f) {
			response.Error(c, 10004, "password error!")
		} else {
			c.SetCookie("auth", token, 1000, "/", "127.0.0.1:8080", false, true)
			response.OkWithData(c, "login successful!")
			G_username = f.Username
		}
	}
}

func PasswdIsOk(f config.LoginForm) bool {
	passwd, _ := dao.G_client.HGet("users", f.Username).Result()
	return passwd == f.Password
}

func IsLogin(c *gin.Context) bool {
	token, err := c.Cookie("auth")
	return TokenIsOk(token) && err == nil
}

func TokenIsOk(token string) bool {
	j := jwts.NewJwt()
	_, err := j.Check(token, "redrock")
	return err == nil
}
