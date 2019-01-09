package controller_auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"hash"
	"io"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-server/controller/http_common"
	"lemon-robot-server/model"
	"lemon-robot-server/service/service_auth"
	"lemon-robot-server/sysinfo"
)

var HashObj hash.Hash

func authLrc(ctx *gin.Context) {
	reqAuthLrc := model.ReqAuthLrc{}
	err := ctx.BindJSON(&reqAuthLrc)
	if err != nil {
		logger.Error("err", err)
	}
	http_common.Success(ctx, service_auth.GenerateJwtTokenStr(reqAuthLrc.Lrct))
}

func calculateLrcps(lrcp string) string {
	if HashObj == nil {
		HashObj = hmac.New(sha256.New, sysinfo.GetHmacKeyBytes())
	}
	HashObj.Reset()
	io.WriteString(HashObj, lrcp)
	return fmt.Sprintf("%x", HashObj.Sum(nil))
}
