package database

import (
	"log"
	"os"

	"github.com/sajid414/Go-react-auth/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("error in database connection")
	}
	log.Println("database connection successful")

	db.AutoMigrate(new(model.User))
	DBConn = db

}
