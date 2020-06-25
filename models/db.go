package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	db := DbConn()
	//Database migration
	// Creates table in database
	db.Debug().AutoMigrate(&Products{})
	db.Debug().AutoMigrate(&Auction{})
	log.Println()

}

func DbConn() (db *gorm.DB) {
	//Load .env file
	// e := godotenv.Load()
	// if e != nil {
	// 	fmt.Print(e)
	// }
	db_name := "bid_db"
	db_pass := "Bunty12345!"
	db_user := "postgres"
	db_host := "localhost"
	// db_port := 5434
	//dbStatus := os.Getenv("app_status")
	//Build connection string
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", db_host, db_user, db_name, db_pass)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	return
}
