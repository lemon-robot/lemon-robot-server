package service_lrc

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lrustring"
	"lemon-robot-server/db"
	"lemon-robot-server/entity"
	"lemon-robot-server/sysinfo"
)

var HashObj hash.Hash

func GenerateLrc(lrcp string) entity.Lrc {
	lrcEntity := entity.Lrc{}
	lrcEntity.LrcKey = lrustring.Uuid()
	lrcEntity.Lrct = lrustring.Uuid()
	lrcEntity.LrcpSecret = CalculateLrcps(lrcp)
	db.Db().Create(&lrcEntity)
	return lrcEntity
}

func CheckLrc(lrct, lrcp string) bool {
	lrcEntity := entity.Lrc{}
	db.Db().First(&lrcEntity, &entity.Lrc{Lrct: lrct})
	if lrcEntity.ID == 0 {
		return false
	}
	return lrcEntity.LrcpSecret == CalculateLrcps(lrcp)
}

func CalculateLrcps(lrcp string) string {
	if HashObj == nil {
		HashObj = hmac.New(sha256.New, sysinfo.GetHmacKeyBytes())
	}
	HashObj.Reset()
	io.WriteString(HashObj, lrcp)
	return fmt.Sprintf("%x", HashObj.Sum(nil))
}

// lrc self repair,
// If the LRC table in the database is empty,
// then an LRC is automatically created randomly and the information is displayed in the console.
func SelfRepair() {
	logger.Info("Start self-repair: LRC")
	//count := db.Db().Count(&entity.Lrc{})
	var count int
	db.Db().Model(&entity.Lrc{}).Count(&count)
	if count == 0 {
		lrcpNew := lrustring.Uuid()
		lrcNew := GenerateLrc(lrcpNew)
		logger.Warn("The number of LRCs is 0, and an LRC is automatically generated.")
		logger.Warn("LRCT: " + lrcNew.Lrct)
		logger.Warn("LRCP: " + lrcpNew)
	}
	logger.Info("Self-repair completed: LRC")
}
