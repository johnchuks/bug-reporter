package main

import (
	"log"
	"os"

	"github.com/johnchuks/feature-reporter/controllers"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/johnchuks/feature-reporter/slackapi"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file:", err)
	}
	a := controllers.App{
		SlackVerificationToken: os.Getenv("SLACK_TOKEN"),
	}
	a.Initialize(
		"127.0.0.1",
		"5432",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	client := slack.New(os.Getenv("BOT_TOKEN"))
	slack := &slackapi.SlackListener{
		Client: client,
		BotID: os.Getenv("BOT_ID"),
	}

	// Listen for incoming slack events
	go slack.ListenAndResponse()

	a.Run(":9000")
}
