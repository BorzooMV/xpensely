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

	dbWithPostgresUser := services.Database{Name: "postgres"}.ConnectDb()
	defer dbWithPostgresUser.Close()

	dbWithApplicationUser := services.Database{Name: os.Getenv("POSTGRES_DB_NAME")}.ConnectDb()
	defer dbWithApplicationUser.Close()

	createDbQs := fmt.Sprintf("CREATE DATABASE %v;", os.Getenv("POSTGRES_DB_NAME"))
	_, err = dbWithPostgresUser.Exec(createDbQs)
	if err != nil {
		fmt.Printf("Can't create database:%v\n", err.Error())
	} else {
		fmt.Printf("Database %s created successfully.\n", os.Getenv("POSTGRES_DB_NAME"))
	}

	// Create Users table
	createUsersTableQs := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		firstname VARCHAR(50) NOT NULL,
		lastname VARCHAR(50) NOT NULL,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(100) NOT NULL,
		is_admin BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := dbWithApplicationUser.Exec(createUsersTableQs); err != nil {
		fmt.Printf("Can't create users table:%v\n", err)
	} else {
		fmt.Println("Table users created successfully")
	}
}
