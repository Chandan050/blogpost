package database

import (
	"fmt"
	"log"

	"github.com/Chandan050/blogwebsite/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "root"
	password = "050220"
	hostname = "127.0.0.1:3306"
	dbname   = "blogWebsite"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

var DB *gorm.DB

func DbConnection() {
	fmt.Println("connection started")

	// Open the database connection
	db, err := gorm.Open(mysql.Open(dsn("")), &gorm.Config{})
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}

	// Explicitly select the database
	result := db.Exec("USE " + dbname)
	if result.Error != nil {
		log.Printf("Error %s when selecting database\n", result.Error)
		return
	}

	// Connection successful
	log.Printf("Connected to DB %s successfully\n", dbname)
	fmt.Println("connection ended")

	// Assign the database instance
	DB = db

	// AutoMigrate after selecting the database
	DB.AutoMigrate(&models.User{}, &models.Blog{})
}
