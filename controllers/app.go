package controllers

import (
	"net/http"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/johnchuks/bug-reporter/models"
	"github.com/johnchuks/bug-reporter/responses"
)

type App struct {
	Router *mux.Router
	DB *gorm.DB
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
	a.Router.Use()
	a.Router.HandleFunc("/", home).Methods("GET")
}

func (a *App) Run(port string) {
	log.Printf("\nServer starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func home(w http.ResponseWriter, r *http.Request) { // this is the home route
	responses.JSON(w, http.StatusOK, "Welcome To Bug Reporter service")
}