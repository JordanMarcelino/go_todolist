package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:wkwkland@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	return db
}
