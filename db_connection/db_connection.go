package db_connection

import (
	"fmt"
)
import "os"
import "github.com/jinzhu/gorm"
import _ "github.com/lib/pq"

var db *gorm.DB

func ConnectDB() (*gorm.DB, error){
	dialect := os.Getenv("DIALECT")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbUri := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password,
		dbPort)

	fmt.Println(dbUri)

	db, err := gorm.Open(dialect, dbUri)

	return db, err
}
