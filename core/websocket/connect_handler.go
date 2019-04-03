package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/service/service_auth"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connectPool = make(map[string]*websocket.Conn)

func dealConnectionClose(conCode string) {
	delete(connectPool, conCode)
	logger.Info(fmt.Sprint("A Websocket connection is broken, code: ", conCode))
	logger.Info(fmt.Sprint("The number of remaining connections on the current node: ", len(connectPool)))
}

func ConnectHandler(context *gin.Context) {
	token := context.Param("token")
	fmt.Println()
	if service_auth.CheckToken(token) {
		os := context.Param("os")
		arch := context.Param("arch")
		dispatcherVersion := context.Param("clientVersion")
		logger.Info(fmt.Sprintf("Websocket was successfully established! Client OS: %v, ARCH: %v, Dispatcher Version: %v", os, arch, dispatcherVersion))
	} else {
		logger.Warn("An illegal websocket connection request has been rejected by the system. token: " + token)
		return
	}
	c, err := upGrader.Upgrade(context.Writer, context.Request, nil)
	if c == nil || err != nil {
		logger.Error("Unable to create websocket connection properly", err)
		return
	}
	conCode := lrustring.Uuid()
	connectPool[conCode] = c
	logger.Info(fmt.Sprint("Websocket connection is successfully established! Code: ", conCode))
	logger.Info(fmt.Sprint("The number of remaining connections on the current node: ", len(connectPool)))
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				dealConnectionClose(conCode)
			default:
				logger.Error("Errors occurred while reading cancelled messages from websocket", err)
			}
			break
		}
		log.Printf("Receive messages from websocket: %s", message)
	}
}
