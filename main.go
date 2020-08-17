package main

import (
	"log"
	"os"

	"github.com/johnchuks/feature-reporter/controllers"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file:", err)
	}
	client := slack.New(os.Getenv("BOT_TOKEN"))
	a := &controllers.App{
		SlackVerificationToken: os.Getenv("SLACK_TOKEN"),
		SlackClient: client,
	}
	a.Initialize(
		"127.0.0.1",
		"5432",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	a.Run(":9000")
}
