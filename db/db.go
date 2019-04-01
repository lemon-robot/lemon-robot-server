package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lemon-robot-server/entity"

	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-server/sysinfo"
	"os"
)

var DbObj *gorm.DB

func InitDb() *gorm.DB {
	// set table name rule
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return sysinfo.LrConfig().DbTablePrefix + defaultTableName
	}
	logger.Info("The system started trying to connect to the database.")
	db, err := gorm.Open(sysinfo.LrConfig().DbType, sysinfo.LrConfig().DbUrl)
	// if enabled debug mode, show gorm log
	db.LogMode(sysinfo.LrConfig().DebugMode)
	DbObj = db
	if err != nil {
		logger.Error("The system could not continue to run because it could not establish a connection with the database", err)
		os.Exit(1)
	}
	updateDb()
	logger.Info("Database connection completed!")
	return DbObj
}

func updateDb() {
	DbObj.AutoMigrate(
		entity.Config{},
		entity.DispatcherInstance{},
		entity.FileResource{},
		entity.Namespace{},
		entity.ParamFileDefine{},
		entity.ParamValueDefine{},
		entity.Task{},
		entity.User{})
}

func Db() *gorm.DB {
	if DbObj == nil {
		InitDb()
	}
	return DbObj
}
