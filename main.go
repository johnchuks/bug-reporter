package main

import (
	"os"
	"fmt"
	"log"
	"github.com/johnchuks/bug-reporter/controllers"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(os.Getenv("DB_NAME"))

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file:", err)
	}
	a := controllers.App{}
	a.Initialize(
		"127.0.0.1",
		"5432",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	a.Run(":9000")
}