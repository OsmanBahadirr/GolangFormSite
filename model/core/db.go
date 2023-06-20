package core

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() *sql.DB {
	if err := godotenv.Load("config/.env"); err != nil {
		fmt.Printf("Error while loading env file: %s\n", err.Error())
		return DB
	}

	connInfo := connection{
		Host:     os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	connString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		connInfo.Host,
		connInfo.Port,
		connInfo.User,
		connInfo.Password,
		connInfo.DBName,
	)
	var err error
	DB, err = sql.Open("postgres", connString)

	if err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err.Error())
	} else {
		fmt.Println("Databse is opened")
	}

	err = DB.Ping()

	if err != nil {
		fmt.Printf("Error: Couldnot ping database: %s\n", err.Error())
		return DB
	} else {
		fmt.Println("Database pinged successfully")
	}
	return DB
}
