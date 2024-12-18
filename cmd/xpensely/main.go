package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BorzooMV/xpensely/internal"
	"github.com/BorzooMV/xpensely/services"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error reading environment variables: %v", err)
	}

	// Connect to DB
	db := services.Database{Name: os.Getenv("POSTGRES_DB_NAME")}.ConnectDb()
	defer db.Close()

	// Create server
	server := internal.CreateEchoServer(db)

	// Start the server
	if err := server.Start(fmt.Sprintf(":%v", os.Getenv("SERVER_LISTENING_PORT"))); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
