package db

import (
	"external_api/config"
	// "ECHO-GORM/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connect_string)
	// defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic("DB Connection Error")
	}
}

func DbManager() *gorm.DB {
	return db
}
