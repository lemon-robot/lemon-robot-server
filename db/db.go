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
		return sysinfo.LrServerConfig().DbTablePrefix + defaultTableName
	}
	logger.Info("The system started trying to connect to the database.")
	db, err := gorm.Open(sysinfo.LrServerConfig().DbType, sysinfo.LrServerConfig().DbUrl)
	// if enabled debug mode, show gorm log
	db.LogMode(sysinfo.LrServerConfig().DebugMode)
	DbObj = db
	db.Set("gorm:auto_preload", true)
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
		entity.DispatcherMachine{},
		entity.DispatcherOnline{},
		entity.DispatcherTag{},
		entity.FileResource{},
		entity.Namespace{},
		entity.ParamFileDefine{},
		entity.ParamValueDefine{},
		entity.ServerNode{},
		entity.Task{},
		entity.User{},
		entity.EnvironmentComponent{},
		entity.EnvironmentComponentVersion{},
		entity.EnvironmentComponentDependencyRelation{},
		entity.OperatePlatform{})
}

func Db() *gorm.DB {
	if DbObj == nil {
		InitDb()
	}
	return DbObj
}
