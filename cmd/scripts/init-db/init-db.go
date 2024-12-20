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

	// Create database
	createDbQs := fmt.Sprintf("CREATE DATABASE %v;", os.Getenv("POSTGRES_DB_NAME"))
	_, err = dbWithPostgresUser.Exec(createDbQs)
	if err != nil {
		fmt.Printf("Can't create database:%v\n", err.Error())
	} else {
		fmt.Printf("Database %s created successfully.\n", os.Getenv("POSTGRES_DB_NAME"))
	}

	// A function to auto update updated_at field
	addAutoUpdateTriggerQs := `
	CREATE OR REPLACE FUNCTION update_updated_at_on_update()
	RETURNS TRIGGER
	LANGUAGE plpgsql
	AS $$
	BEGIN
		NEW.updated_at = CURRENT_TIMESTAMP;
		RETURN NEW;
	END;	
	$$;
	`
	if _, err := dbWithApplicationUser.Exec(addAutoUpdateTriggerQs); err != nil {
		log.Fatalf("Can't create auto update function:%v\n", err)
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

	// Add trigger for users to update automatically
	addTriggerForUsersQs := `
	CREATE OR REPLACE TRIGGER update_updated_at
	BEFORE UPDATE ON users
	FOR EACH ROW
	EXECUTE FUNCTION update_updated_at_on_update()
	`
	if _, err := dbWithApplicationUser.Exec(addTriggerForUsersQs); err != nil {
		log.Fatalf("Can't create trigger for users table:%v\n", err)
	}

	// Create Expenses table
	createExpensesTableQs := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount DECIMAL(10,2) NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	if _, err := dbWithApplicationUser.Exec(createExpensesTableQs); err != nil {
		fmt.Printf("Can't create expenses table:%v\n", err)
	} else {
		fmt.Println("Table expenses created successfully")
	}

	// Add trigger for expenses to update automatically
	addTriggerForExpensesQs := `
	CREATE OR REPLACE TRIGGER update_updated_at
	BEFORE UPDATE ON expenses
	FOR EACH ROW
	EXECUTE FUNCTION update_updated_at_on_update()
	`
	if _, err := dbWithApplicationUser.Exec(addTriggerForExpensesQs); err != nil {
		log.Fatalf("Can't create trigger for expenses table:%v\n", err)
	}
}
