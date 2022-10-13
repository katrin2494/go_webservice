package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type SqlHandler struct {
	Db *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := "upgrade_user:123456@tcp(127.0.0.1:3306)/upgrade?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	sqlHandler := &SqlHandler{}
	sqlHandler.Db = db

	return sqlHandler
}
