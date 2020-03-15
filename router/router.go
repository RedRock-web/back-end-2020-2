package router

import (
	"back-end-2020-1/app/account"
	"back-end-2020-1/app/game"
	"back-end-2020-1/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", account.Register)
	r.POST("/login", account.Login)

	g := r.Group("/", middleware.AuthCheck())
	{
		g.PUT("/enter", game.Enter)
		g.PUT("/retire", game.Retire)
		g.GET("/leader_board")
	}
}
