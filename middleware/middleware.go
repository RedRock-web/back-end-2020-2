package middleware

import (
	"back-end-2020-1/app/account"
	"back-end-2020-1/app/jwts"
	"back-end-2020-1/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckHaveEnterGame() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth")
		if err != nil {
			response.Error(c, 10009, "needed login!")
			c.Abort()
		} else {
			j := jwts.NewJwt()
			fmt.Println(token)
			f, _ := j.Check(token, "redrock")
			account.G_username = f.Username
			c.Next()
		}
	}
}
