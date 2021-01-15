package configs

import (
	"fmt"
	"github.com/maulanay85/go-ecommerce-product/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	props, err := ReadPropertiesFile(AbsPath + "/application.properties")
	if err != nil {
		log.Fatal(err)
	}

	url := props["database.url"]
	port := props["database.port"]
	username := props["database.username"]
	password := props["database.password"]
	name := props["database.name"]

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", url, username, password, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Item{})
	DB = db
}
