package storage

import (
	"fmt"
	"os"

	"github.com/ebikode/peaq-challenge/challenge3/exchange/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // For Mysql Dialects import
)

// MDatabase ...
type MDatabase struct {
	db *gorm.DB
}

// Config ...
type Config struct {
}

// New ...
func New() *Config {
	return &Config{}
}

// InitDB ..
func (config *Config) InitDB() (*MDatabase, error) {

	var err error
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	charset := "utf8mb4_unicode_ci"

	dbHost = "127.0.0.1"
	dbPort = "3306"
	dbName = "peaq_analytics_goDB"
	dbUser = "root"
	dbPass = ""

	// First Connect to create a DB if it doesn't exist
	conHostURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True", dbUser, dbPass, dbHost, dbPort) //Build connection string
	cdb, err := gorm.Open("mysql", conHostURI)
	defer cdb.Close()

	if err != nil {
		return nil, err
	}

	createQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE %s;", dbName, charset)
	cErr := cdb.Exec(createQuery).Error
	if cErr != nil {
		fmt.Println(cErr.Error())
		return nil, cErr
	}

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", dbUser, dbPass, dbHost, dbPort, dbName) //Build connection string
	fmt.Println(dbURI)

	mdb := new(MDatabase)

	mdb.db, err = gorm.Open("mysql", dbURI)
	if err != nil {
		fmt.Print(err)
		return mdb, err
	}

	// Migrating tables to database
	mdb.db.Debug().AutoMigrate(
		&models.Rate{},
		&models.GrowthRecord{},
	) //Database migration

	// defer mdb.db.Close()

	return mdb, nil

}
