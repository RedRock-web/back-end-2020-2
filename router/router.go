package router

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) {
	r.POST("/register")
	r.POST("/login")
	r.POST("/enter")
	r.POST("/retire")
	r.GET("/leader_board")
}
