package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BorzooMV/xpensely/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := services.Database{Name: os.Getenv("POSTGRES_DB_NAME")}.ConnectDb()
	defer db.Close()

	fmt.Println("Start cleaning database...")
	db.Exec("DELETE FROM users;")
	db.Exec("DELETE FROM expenses;")
	fmt.Println("Database is empty")
	db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
	db.Exec("ALTER SEQUENCE expenses_id_seq RESTART WITH 1")
}
