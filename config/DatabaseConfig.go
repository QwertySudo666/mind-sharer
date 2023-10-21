package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var db *gorm.DB

func ConfigureMySQLConnection() (*gorm.DB, error) {
	dsn := "root:localhost@tcp(127.0.0.1:3306)/mind_sharer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db, nil
}
