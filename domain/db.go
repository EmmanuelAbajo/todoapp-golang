package domain

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func GetDbConnectionURL() string {
	_ = godotenv.Load("conf.env")

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",username,password,host,port,name)
	fmt.Printf("DbURL: %s\n", dbURL)

	return dbURL
}
