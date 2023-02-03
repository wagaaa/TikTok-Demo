package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"tiktok.service/biz/utils"
)

var DB *gorm.DB
var cfg = utils.Cfg.Database

// Init 初始化数据库
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Passwd,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
