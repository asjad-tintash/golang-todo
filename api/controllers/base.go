package controllers

import (
	"fmt"
	"github.com/asjad-tintash/golang-todo/api/middlewares"
	"github.com/asjad-tintash/golang-todo/api/responses"
	"github.com/asjad-tintash/golang-todo/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	Db     *gorm.DB
}

func (a *App) Initialize(DbHost, DbName, DbUser, DbPassword, DbPort string) {
	DatabaseURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", DbHost, DbPort, DbName, DbUser, DbPassword)
	var err error
	a.Db, err = gorm.Open("postgres", DatabaseURI)
	if err != nil {
		fmt.Printf("Unable to connect to the database asjad\n")
		log.Fatal(err)
	} else {
		fmt.Printf("Connected to the database\n")
	}

	a.Db.Debug().AutoMigrate(&models.User{}, &models.Task{}, &models.UnAssignedTask{})
	a.Db.Debug().Model(&models.Task{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	a.Db.Debug().Model(&models.UnAssignedTask{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	a.Router = mux.NewRouter().StrictSlash(true)

	a.InitializeRoutes()
}

func (a *App) InitializeRoutes() {

	a.Router.Use(middlewares.SetContentTypeMiddleware)
	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/register", a.Register).Methods("POST")
	a.Router.HandleFunc("/login", a.Login).Methods("POST")

	s := a.Router.PathPrefix("/api").Subrouter()
	s.Use(middlewares.AuthJwtVerify)

	s.HandleFunc("/user/{id:[0-9]+}", a.DeleteUser).Methods("DELETE")
	s.HandleFunc("/task", a.CreateTask).Methods("POST")
	s.HandleFunc("/task", a.GetTasks).Methods("GET")
	s.HandleFunc("/task/{id:[0-9]+}", a.UpdateTask).Methods("PATCH")
	s.HandleFunc("/task/{id:[0-9]+}", a.DeleteTask).Methods("DELETE")

	s.HandleFunc("/assign_task", a.AssignTask).Methods("POST")
}

func (a *App) RunServer() {
	log.Printf("Server starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to the app")
}
