package storage

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var database *sql.DB

func init() {
	config := mysql.NewConfig()
	config.Addr = os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	config.DBName = os.Getenv("DB_DATABASE")
	config.User = os.Getenv("DB_USERNAME")
	config.Passwd = os.Getenv("DB_PASSWORD")
	// TODO: check
	config.Timeout = 30_000_000_000      // 30 seconds
	config.ReadTimeout = 10_000_000_000  // 10 seconds
	config.WriteTimeout = 10_000_000_000 // 10 seconds
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	database = db
	var version string
	database.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to database:", version)
}

func GetSQL() *sql.DB {
	return database
}

func Close() {
	fmt.Println("Closing database connection")
	err := database.Close()
	if err != nil {
		panic(err)
	}
}
