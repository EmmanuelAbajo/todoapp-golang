package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var dbConn *sql.DB
var err error

func init() {
	_ = godotenv.Load("conf.env")

	log.Println("Conecting to Database...")
	dbConn, err = sql.Open("mysql", GetDbConnectionURL())
	if err != nil {
		panic(err.Error())
	}
	log.Println("Connection to Database successful!")
}

func GetDbConnectionURL() string {
	_ = godotenv.Load("conf.env")

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, name)
	fmt.Printf("DbURL: %s\n", dbURL)

	return dbURL
}

func Connection() *sql.DB {
	return dbConn
}
