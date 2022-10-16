package database

import (
	"log"

	"github.com/Justincletus003/go-prod-ord/models"
	"gorm.io/driver/mysql"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

var DB *gorm.DB
// type DbInstance struct {
// 	Db *gorm.DB
// }
// var Database DbInstance

func Connect()  {
	db, err := gorm.Open(mysql.Open("root:Pass#123@tcp(localhost:3306)/prod_ord_db?parseTime=true"), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("Database connection falied!")
	}
	
	log.Println("Database connected!")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	DB = db
	// Database = DbInstance{
	// 	Db: db,
	// }
}

