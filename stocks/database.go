package stocks

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:razeem19@tcp(host.docker.internal:3306)/stock_market"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = DB.AutoMigrate(
		&Stock{},
		&Order{},
	)
	if err != nil {
		log.Printf("Failed to auto-migrate the database: %v", err)
		return
	}

	log.Println("Database connection and migration successful")
}
