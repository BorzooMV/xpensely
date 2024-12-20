package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/BorzooMV/xpensely/services"
	"github.com/joho/godotenv"
)

func main() {
	var Users []struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Password  string `json:"password"`
		Username  string `json:"username"`
		Email     string `json:"email"`
	}

	var Expenses []struct {
		UserId      int     `json:"user_id"`
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read users sample file
	usersFile, err := os.ReadFile("assets/data/sample-users.json")
	if err != nil {
		log.Fatalf("Couldn't read the recipe file:\n%v\n", err.Error())
	}

	err = json.Unmarshal(usersFile, &Users)
	if err != nil {
		log.Fatalf("Couldn't unmarshal the recipe file content:\n%v\n", err.Error())
	}

	// Read expenses sample file
	expensesFile, err := os.ReadFile("assets/data/sample-expenses.json")
	if err != nil {
		log.Fatalf("Couldn't read the recipe file:\n%v\n", err.Error())
	}

	err = json.Unmarshal(expensesFile, &Expenses)
	if err != nil {
		log.Fatalf("Couldn't unmarshal the recipe file content:\n%v\n", err.Error())
	}

	db := services.Database{Name: os.Getenv("POSTGRES_DB_NAME")}.ConnectDb()
	defer db.Close()

	// Start seeding
	fmt.Println("Start seeding database with fake data...")

	// Seeding users
	for _, item := range Users {
		qs := "INSERT INTO users (firstname, lastname, username, email, password) VALUES ($1,$2,$3,$4,$5);"
		_, err := db.Exec(qs, item.Firstname, item.Lastname, item.Username, item.Email, item.Password)
		if err != nil {
			log.Fatalf("Couldn't insert into db:\n%v\n", err.Error())
		}
	}

	// Seeding expenses
	for _, item := range Expenses {
		qs := "INSERT INTO expenses (user_id, amount, description) VALUES ($1,$2,$3);"
		_, err := db.Exec(qs, item.UserId, item.Amount, item.Description)
		if err != nil {
			log.Fatalf("Couldn't insert into db:\n%v\n", err.Error())
		}
	}

	fmt.Println("Seeding completed")
}
