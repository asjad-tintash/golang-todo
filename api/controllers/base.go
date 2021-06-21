package controllers

import (
	"../../models"
	"../middlewares"
	"../responses"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Db *gorm.DB
}

func (a *App) Initialize(DbHost, DbName, DbUser, DbPassword, DbPort string){
	DatabaseURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", DbHost, DbPort, DbName, DbUser, DbPassword)
	var err error
	a.Db, err = gorm.Open("postgres", DatabaseURI)
	if err != nil {
		fmt.Printf("Unable to connect to the database\n")
		log.Fatal(err)
	} else {
		fmt.Printf("Connected to the database\n")
	}

	a.Db.Debug().AutoMigrate(&models.User{})
	a.Router = mux.NewRouter().StrictSlash(true)

	a.InitializeRoutes()
}

func (a *App) InitializeRoutes() {

	a.Router.Use(middlewares.SetContentTypeMiddleware)
	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/register", a.Register).Methods("POST")
	a.Router.HandleFunc("/login", a.Login).Methods("POST")
}

func (a *App) RunServer() {
	log.Printf("Server starting on port 5000")
	log.Fatal(http.ListenAndServe(":5000", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to the app")
}