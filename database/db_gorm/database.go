package db_gorm

import (
	"api/config/setting"
	"api/model"
	"api/util/log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func OpenConnect() {
	var err error
	var config *gorm.Config
	if gin.Mode() == gin.ReleaseMode {
		config = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction: true,
		}
	}else {
		config = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction: true,
			Logger:logger.Default.LogMode(logger.Info),
		}
	}

	dsn := setting.DataBaseConfig.UserName + ":" + setting.DataBaseConfig.Password + "@tcp(" + setting.DataBaseConfig.Host + ":" + setting.DataBaseConfig.Port + ")/" + setting.DataBaseConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), config)

	if err != nil {
		log.Error(err)
	}

	if err != nil {
		log.Error(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	model.SetDB(db)
}

func GetDB() *gorm.DB {
	return db
}
