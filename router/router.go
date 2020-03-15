package router

import (
	"back-end-2020-1/app/account"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", account.Register)
	r.POST("/login")
	r.POST("/enter")
	r.POST("/retire")
	r.GET("/leader_board")
}
