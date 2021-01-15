package configs

import (
	"github.com/jinzhu/gorm"
	"github.com/maulana/cari-jasa/models"
)

var DB *gorm.DB

func ConnectToDB() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=carijasa password=root")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(models.CjUser{})
	db.AutoMigrate(models.CjJob{})
	DB = db
}
