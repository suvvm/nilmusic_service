package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"suvvm.work/nilmusic_service/config"
)

var (
	DB *gorm.DB
)
// InitDB 初始化数据库链接
//
// 入参
//	dbConfig *config.MysqlConfig 	// db配置信息
// 返回
//	error	// 错误信息
func InitDB(dbConfig *config.MysqlConfig) error {
	// 构造DSN
	dsn := dbConfig.GetDSN()
	// 根据DSN获取db链接
	gDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("init gorm db conn err=%s", err)
		return err
	}
	DB = gDB
	db , err := DB.DB()	// 获取sql.DB
	if err != nil {
		log.Printf("get sql.DB conn err=%s", err)
		return err
	}
	db.SetMaxIdleConns(dbConfig.MaxIdle)	// 设置链接池最大空闲链接数
	db.SetMaxOpenConns(dbConfig.MaxOpen)	// 设置数据库最大打开链接数
	log.Printf("init db conn sucess")
	return nil
}