package middleware

import (
	"back-end-2020-1/app/game"
	"back-end-2020-1/response"
	"github.com/gin-gonic/gin"
)

func CheckHaveEnterGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		if game.HaveEnter() {
			c.Next()
		} else {
			response.Error(c, 10011, "needed enter game!")
			c.Abort()
		}
	}
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("auth")
		if err != nil {
			response.Error(c, 10009, "needed login!")
			c.Abort()
		} else {
			//TODO: token 的 check 有 bug，暂时无法获取 username 和 password
			//j := jwts.NewJwt()
			//f, _ := j.Check(token, "redrock")
			//account.G_username = f.Username
			c.Next()
		}
	}
}
