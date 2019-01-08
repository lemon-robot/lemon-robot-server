package controller_lrc

import "github.com/gin-gonic/gin"

func RegApis(engine *gin.Engine) {
	engine.GET("/", list)
}

func list(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "liuri",
	})
}
