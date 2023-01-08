package config

import (
	"fmt"
	"github.com/ilyamudritskiy/bookstore_management/pkg/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	var db_credentials = fmt.Sprintf(
		"%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.DB_USERNAME,
		settings.DB_PASSWD,
		settings.DB_TABLENAME)

	db, err := gorm.Open("mysql", db_credentials)

	if err != nil {
		panic(err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
