package account

import (
	"back-end-2020-1/app/jwts"
	"back-end-2020-1/config"
	"back-end-2020-1/dao/dao_mysql"
	"back-end-2020-1/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	f := config.LoginForm{}

	if err := c.ShouldBindJSON(&f); err != nil {
		errors.New("bind json error!")
		response.FormError(c)
	} else if IsRegiste(f.Username) {
		response.Error(c, 10003, "user exist!")
	} else {
		dao_mysql.Insert(dao_mysql.User{Username: f.Username, Password: f.Password}, "register insert record error!")
		token := GetJwt(f, "register create jwt error!")
		c.SetCookie("auth", token, 1000, "/", "127.0.0.1", false, true)
	}
}

func IsRegiste(username string) bool {
	var user dao_mysql.User
	dao_mysql.G_db.Where("username = ?", username).First(&user)
	return user.ID != 0
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

	if err := c.ShouldBindJSON(&f); err != nil {
		errors.New("bind json error!")
		response.FormError(c)
	} else if IsLogin(c) {
		response.Ok(c)
	} else {
		if !IsRegiste(f.Username) {
			response.Error(c, 10005, "unregistered!")
		} else if !PasswdIsOk(f) {
			response.Error(c, 10004, "password error!")
		} else {
			response.Ok(c)
			c.SetCookie("auth", token, 1000, "/", "127.0.0.1", false, true)
		}
	}
}

func PasswdIsOk(f config.LoginForm) bool {
	var user dao_mysql.User
	dao_mysql.G_db.Where(dao_mysql.User{
		Username: f.Username,
		Password: f.Password,
	}).First(&user)
	return user.ID != 0
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
