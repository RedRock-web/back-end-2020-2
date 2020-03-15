package middleware

import (
	"back-end-2020-1/response"
	"github.com/gin-gonic/gin"
)

func CheckHaveEnterGame() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("auth")
		if err != nil {
			response.Error(c, 10009, "needed login!")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
