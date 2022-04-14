package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyDB struct {
	dbNameString string
	db           *gorm.DB
}

func NewMyDB(name string) (*MyDB, error) {
	dsn := "root:84490979@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MyDB{db: db}, nil
}
