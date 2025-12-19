package db

import (
	"basketball/conf"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb(mysqlConf *conf.Db) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("fatal error connect db: %w", err))
	}
	return db
}
