package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BorzooMV/xpensely/internal"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error reading environment variables: %v", err)
	}

	server := internal.CreateEchoServer()

	if err := server.Start(fmt.Sprintf(":%v", os.Getenv("SERVER_LISTENING_PORT"))); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
