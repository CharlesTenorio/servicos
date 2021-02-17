package config

import (
	"fmt"

	"github.com/CharlesTenorio/servicos/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err := gorm.Open(postgres.Open(connect_string), &gorm.Config{})
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&model.Time{})

}

func DbManager() *gorm.DB {
	return db
}
