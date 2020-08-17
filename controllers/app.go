package controllers

import (
	"github.com/johnchuks/feature-reporter/middlewares"
	"net/http"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/johnchuks/feature-reporter/models"
	"github.com/johnchuks/feature-reporter/responses"
	"github.com/slack-go/slack"
)

type App struct {
	Router *mux.Router
	DB *gorm.DB
	SlackVerificationToken string
	SlackClient *slack.Client
}


func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", 
		host, port, user, dbname, password)

	var err error

	a.DB, err = gorm.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("An error occurred:", err)
	} else {
		log.Printf("Successfully connected to database %s", dbname)
	}

	a.DB.Debug().AutoMigrate(&models.Report{}) // database migration
	
	a.Router = mux.NewRouter().StrictSlash(true)
	a.intializeRoutes()

}

func (a *App) intializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware)
	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/event/feature", a.SlackHandler).Methods("POST")
}

func (a *App) Run(port string) {
	log.Printf("\nServer starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func home(w http.ResponseWriter, r *http.Request) { // this is the home route
	responses.JSON(w, http.StatusOK, "Welcome To Bug Reporter service")
}