package database

import (
	"entry_task/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyDB struct {
	dbNameString string

	db *gorm.DB
	//mongodb
}

func NewMyDB(conf *config.Mysql) (*MyDB, error) {
	//to use config
	dsn := "root:84490979@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	mysqlDB, err := gormdb.DB()
	if err != nil {
		return nil, err
	}
	mysqlDB.SetMaxIdleConns(1000)
	return &MyDB{db: gormdb}, nil
}
