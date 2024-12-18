package services

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Database struct {
	Name string
}

func (d Database) ConnectDb() *sql.DB {
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port, _ := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 0)
	user := os.Getenv("POSTGRES_USER")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, d.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}
