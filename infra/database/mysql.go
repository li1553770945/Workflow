package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"workflow_http/infra/conf"
)

func NewMySQL(cfg *conf.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MysqlConfig.Username, cfg.MysqlConfig.Password, cfg.MysqlConfig.Address, cfg.MysqlConfig.Database)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}
