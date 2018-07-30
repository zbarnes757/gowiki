package database

import (
	"fmt"
	"gowiki/models"
	"os"

	"github.com/jinzhu/gorm"
	// used to setup the connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// SetupDB will migrate schemas and prepare the database and return a connection
func SetupDB() *gorm.DB {
	addr := getDBAddress()
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Migrate the schema
	db.AutoMigrate(&models.Page{})

	return db
}

func getDBAddress() string {
	if addr := os.Getenv("DATABASE_URL"); addr != "" {
		return addr
	}

	return "postgresql://maxroach@localhost:26257/wiki?sslmode=disable"
}
