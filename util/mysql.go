package util

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/helmimuzkr/belajar-golang-restapi/config"
)

func NewDatabaseConnection(config *config.AppConfig) *sql.DB {
	port := config.Database.Port
	host := config.Database.Host
	username := config.Database.Username
	password := config.Database.Password
	dbName := config.Database.DBName

	// Buat data source
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	// Open connection database
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	// Configuration pooling database
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
