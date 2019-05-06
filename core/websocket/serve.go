package websocket

import (
	"github.com/gin-gonic/gin"
)

func Serve(engine *gin.Engine) {
	engine.GET("/ws/:os/:arch/:dispatcherVersion/:machineSign/:token", ConnectHandler)
}
