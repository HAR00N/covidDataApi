package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

func DbConnection() *sql.DB {
	godotenv.Load(".env")
	dbUrl := os.Getenv("db_url")
	database, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	return database
}
