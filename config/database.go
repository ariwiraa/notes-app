package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ariwiraa/notes-app/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	username := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	driver := os.Getenv("DB_DRIVER")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", username, password, host, port, databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect %s database", driver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the %s database", driver)
	}

	db.Debug().AutoMigrate(&domain.Notes{})

	return db
}
